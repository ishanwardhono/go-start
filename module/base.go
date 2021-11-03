package module

import "context"

type BaseModule interface {
	Execute(ctx context.Context) (interface{}, error)
	Validate(ctx context.Context) error
}
