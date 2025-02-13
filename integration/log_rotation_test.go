//go:build !windows
// +build !windows

package integration

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/go-check/check"
	"github.com/rs/zerolog/log"
	checker "github.com/vdemeester/shakers"
	"traefik/v3/integration/try"
)

const traefikTestAccessLogFileRotated = traefikTestAccessLogFile + ".rotated"

// Log rotation integration test suite.
type LogRotationSuite struct{ BaseSuite }

func (s *LogRotationSuite) SetUpSuite(c *check.C) {
	s.createComposeProject(c, "access_log")
	s.composeUp(c)
}

func (s *LogRotationSuite) TearDownSuite(c *check.C) {
	s.composeDown(c)

	generatedFiles := []string{
		traefikTestLogFile,
		traefikTestAccessLogFile,
		traefikTestAccessLogFileRotated,
	}

	for _, filename := range generatedFiles {
		if err := os.Remove(filename); err != nil {
			log.Warn().Err(err).Send()
		}
	}
}

func (s *LogRotationSuite) TestAccessLogRotation(c *check.C) {
	// Start Traefik
	cmd, display := s.traefikCmd(withConfigFile("fixtures/access_log_config.toml"))
	defer display(c)
	defer displayTraefikLogFile(c, traefikTestLogFile)

	err := cmd.Start()
	c.Assert(err, checker.IsNil)
	defer s.killCmd(cmd)

	// Verify Traefik started ok
	verifyEmptyErrorLog(c, "traefik.log")

	waitForTraefik(c, "server1")

	// Make some requests
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/", nil)
	c.Assert(err, checker.IsNil)
	req.Host = "frontend1.docker.local"

	err = try.Request(req, 500*time.Millisecond, try.StatusCodeIs(http.StatusOK), try.HasBody())
	c.Assert(err, checker.IsNil)

	// Rename access log
	err = os.Rename(traefikTestAccessLogFile, traefikTestAccessLogFileRotated)
	c.Assert(err, checker.IsNil)

	// in the midst of the requests, issue SIGUSR1 signal to server process
	err = cmd.Process.Signal(syscall.SIGUSR1)
	c.Assert(err, checker.IsNil)

	// continue issuing requests
	err = try.Request(req, 500*time.Millisecond, try.StatusCodeIs(http.StatusOK), try.HasBody())
	c.Assert(err, checker.IsNil)
	err = try.Request(req, 500*time.Millisecond, try.StatusCodeIs(http.StatusOK), try.HasBody())
	c.Assert(err, checker.IsNil)

	// Verify access.log.rotated output as expected
	logAccessLogFile(c, traefikTestAccessLogFileRotated)
	lineCount := verifyLogLines(c, traefikTestAccessLogFileRotated, 0, true)
	c.Assert(lineCount, checker.GreaterOrEqualThan, 1)

	// make sure that the access log file is at least created before we do assertions on it
	err = try.Do(1*time.Second, func() error {
		_, err := os.Stat(traefikTestAccessLogFile)
		return err
	})
	c.Assert(err, checker.IsNil, check.Commentf("access log file was not created in time"))

	// Verify access.log output as expected
	logAccessLogFile(c, traefikTestAccessLogFile)
	lineCount = verifyLogLines(c, traefikTestAccessLogFile, lineCount, true)
	c.Assert(lineCount, checker.Equals, 3)

	verifyEmptyErrorLog(c, traefikTestLogFile)
}

func logAccessLogFile(c *check.C, fileName string) {
	output, err := os.ReadFile(fileName)
	c.Assert(err, checker.IsNil)
	c.Logf("Contents of file %s\n%s", fileName, string(output))
}

func verifyEmptyErrorLog(c *check.C, name string) {
	err := try.Do(5*time.Second, func() error {
		traefikLog, e2 := os.ReadFile(name)
		if e2 != nil {
			return e2
		}
		c.Assert(string(traefikLog), checker.HasLen, 0)
		return nil
	})
	c.Assert(err, checker.IsNil)
}

func verifyLogLines(c *check.C, fileName string, countInit int, accessLog bool) int {
	rotated, err := os.Open(fileName)
	c.Assert(err, checker.IsNil)
	rotatedLog := bufio.NewScanner(rotated)
	count := countInit
	for rotatedLog.Scan() {
		line := rotatedLog.Text()
		if accessLog {
			if len(line) > 0 {
				if !strings.Contains(line, "/api/rawdata") {
					CheckAccessLogFormat(c, line, count)
					count++
				}
			}
		} else {
			count++
		}
	}

	return count
}
