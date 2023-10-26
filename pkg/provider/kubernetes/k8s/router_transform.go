package k8s

import (
	"context"

	"traefik/v3/pkg/config/dynamic"
)

type RouterTransform interface {
	Apply(ctx context.Context, rt *dynamic.Router, annotations map[string]string) error
}
