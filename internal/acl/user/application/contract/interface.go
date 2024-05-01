package contract

import "context"

type Command interface {}

type HandlerResult interface{}

type HandlerErr error

type Handler interface {
	Handle(ctx context.Context, cmd Command) (HandlerResult, HandlerErr)
}

type Validator interface {
	Validate(target interface{}) error
}