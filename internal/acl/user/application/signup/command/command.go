package command

import (
	"goedabook/internal/acl/user/application/contract"
	"goedabook/internal/acl/user/domain/valueobj"
)

type SignupCommand struct {
	Username valueobj.Username
}

type SignupCommandOptions func(cmd *SignupCommand) error

func NewSignUpCommand(opts ...SignupCommandOptions) (contract.Command, error) {
	cmd := &SignupCommand{}
	for _, opt := range opts {
		if err := opt(cmd); err != nil {
			return nil, err
		}
	}

	return cmd, nil
}

func WithUserName(username valueobj.Username) SignupCommandOptions {
	return func(cmd *SignupCommand) error {
		cmd.Username = username
		return nil
	}
}
