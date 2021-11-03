package log

import (
	"context"

	"github.com/google/uuid"
)

type ContextContent string

const (
	RequestID ContextContent = "request_id"
	//Add here for another context fields
)

var (
	CtxContents = []ContextContent{RequestID}
)

func ContextStory(parent context.Context) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	return SetCtxRequestID(parent)
}

func SetCtxRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, RequestID, uuid.New())
}

func SetCtxVal(ctx context.Context, key ContextContent, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetCtxContent(ctx context.Context) map[ContextContent]interface{} {
	if ctx == nil {
		return nil
	}
	ctxVal := make(map[ContextContent]interface{})
	for _, val := range CtxContents {
		if v := ctx.Value(val); v != nil {
			ctxVal[val] = v
		}
	}
	return ctxVal
}
