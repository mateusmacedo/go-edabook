package handler

import (
	"context"

	"goedabook/pkg/cqrs"
	"goedabook/pkg/validation"
)

var (
	RequiredFieldErrMsg = "field %s is required"
)


type signupHandler struct {
	validators []validation.Validator
}

type SignupHandlerOption func(*signupHandler) error

func WithValidators(validators ...validation.Validator) SignupHandlerOption {
	return func(h *signupHandler) error {
		h.validators = validators
		return nil
	}
}

func NewSignupHandler(opts ...SignupHandlerOption) cqrs.Handler {
	signupHandler := &signupHandler{}
	for _, opt := range opts {
		opt(signupHandler)
	}

	return signupHandler
}

func (h *signupHandler) Handle(ctx context.Context, cmd cqrs.Command) (cqrs.HandlerResult, cqrs.HandlerErr) {
	var res cqrs.HandlerResult
	var err cqrs.HandlerErr

	for _, v := range h.validators {
		if err = v.Validate(cmd); err != nil {
			return res, err
		}
	}


	return res, err
}