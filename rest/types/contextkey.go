package types

import (
	"context"
	"fmt"
)

type ContextKey[T any] string

func (key ContextKey[T]) WithValue(ctx context.Context, t T) context.Context {
	return context.WithValue(ctx, key, t)
}

func (key ContextKey[T]) Value(ctx context.Context) (T, bool) {
	val, ok := ctx.Value(key).(T)
	return val, ok
}

func (key ContextKey[T]) MustValue(ctx context.Context) T {
	val, ok := ctx.Value(key).(T)
	if !ok {
		panic(fmt.Errorf("ContextKey: missing value for %q", key))
	}

	return val
}
