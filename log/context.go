package log

import (
	"context"

	"github.com/google/uuid"
)

func ContextStory(parent context.Context) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	return SetCtxRequestID(parent)
}

func SetCtxRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, "requestID", uuid.New())
}
