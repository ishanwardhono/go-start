package errors

import (
	"errors"
	"log"
	"runtime"
)

type Errs struct {
	err         error
	Code        int
	message     string
	isRetryable bool
}

func New(args ...interface{}) *Errs {
	err := &Errs{err: errors.New("unknown error")}
	for _, arg := range args {
		switch arg.(type) {
		case string:
			err.err = errors.New(arg.(string))
		case *Errs:
			errcpy := *arg.(*Errs)
			err = &errcpy
		case error:
			err.err = arg.(error)
		case int:
			err.Code = arg.(int)
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.Errs: bad call from %s:%d: %v", file, line, args)
		}
	}
	return err
}

func (e *Errs) Error() string {
	return e.err.Error()
}

func (e *Errs) SetMessage(message string) {
	e.message = message
}

func (e *Errs) GetMessage() string {
	return e.message
}

func (e *Errs) IsRetryable() {
	e.isRetryable = true
}
