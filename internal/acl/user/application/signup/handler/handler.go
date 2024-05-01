package handler

import (
	"context"

	"goedabook/internal/acl/user/application/contract"
)

var (
	RequiredFieldErrMsg = "field %s is required"
)


type signupHandler struct {
	validators []contract.Validator
}

type SignupHandlerOption func(*signupHandler) error

func WithValidators(validators ...contract.Validator) SignupHandlerOption {
	return func(h *signupHandler) error {
		h.validators = validators
		return nil
	}
}

func NewSignupHandler(opts ...SignupHandlerOption) contract.Handler {
	signupHandler := &signupHandler{}
	for _, opt := range opts {
		opt(signupHandler)
	}

	return signupHandler
}

func (h *signupHandler) Handle(ctx context.Context, cmd contract.Command) (contract.HandlerResult, contract.HandlerErr) {
	var res contract.HandlerResult
	var err contract.HandlerErr

	for _, v := range h.validators {
		if err = v.Validate(cmd); err != nil {
			return res, err
		}
	}


	return res, err
}