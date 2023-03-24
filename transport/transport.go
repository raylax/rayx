package transport

import (
	"context"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/expression"
)

type Response expression.EValue

type Transport interface {
	Do(ctx context.Context, request *dsl.Request) (Response, error)
	Close()
}

func ref[T any](v T) *T {
	return &v
}
