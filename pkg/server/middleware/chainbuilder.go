package middleware

import (
	"context"

	"github.com/containous/alice"
	"github.com/rs/zerolog/log"
	"traefik/v3/pkg/metrics"
	"traefik/v3/pkg/middlewares/accesslog"
	"traefik/v3/pkg/middlewares/capture"
	metricsmiddleware "traefik/v3/pkg/middlewares/metrics"
	mTracing "traefik/v3/pkg/middlewares/tracing"
	"traefik/v3/pkg/tracing"
)

// ChainBuilder Creates a middleware chain by entry point. It is used for middlewares that are created almost systematically and that need to be created before all others.
type ChainBuilder struct {
	metricsRegistry        metrics.Registry
	accessLoggerMiddleware *accesslog.Handler
	tracer                 *tracing.Tracing
}

// NewChainBuilder Creates a new ChainBuilder.
func NewChainBuilder(metricsRegistry metrics.Registry, accessLoggerMiddleware *accesslog.Handler, tracer *tracing.Tracing) *ChainBuilder {
	return &ChainBuilder{
		metricsRegistry:        metricsRegistry,
		accessLoggerMiddleware: accessLoggerMiddleware,
		tracer:                 tracer,
	}
}

// Build a middleware chain by entry point.
func (c *ChainBuilder) Build(ctx context.Context, entryPointName string) alice.Chain {
	chain := alice.New()

	if c.accessLoggerMiddleware != nil || c.metricsRegistry != nil && (c.metricsRegistry.IsEpEnabled() || c.metricsRegistry.IsRouterEnabled() || c.metricsRegistry.IsSvcEnabled()) {
		chain = chain.Append(capture.Wrap)
	}

	if c.accessLoggerMiddleware != nil {
		chain = chain.Append(accesslog.WrapHandler(c.accessLoggerMiddleware))
	}

	if c.tracer != nil {
		chain = chain.Append(mTracing.WrapEntryPointHandler(ctx, c.tracer, entryPointName))
	}

	if c.metricsRegistry != nil && c.metricsRegistry.IsEpEnabled() {
		chain = chain.Append(metricsmiddleware.WrapEntryPointHandler(ctx, c.metricsRegistry, entryPointName))
	}

	return chain
}

// Close accessLogger and tracer.
func (c *ChainBuilder) Close() {
	if c.accessLoggerMiddleware != nil {
		if err := c.accessLoggerMiddleware.Close(); err != nil {
			log.Error().Err(err).Msg("Could not close the access log file")
		}
	}

	if c.tracer != nil {
		c.tracer.Close()
	}
}
