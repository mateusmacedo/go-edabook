package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"goedabook/internal/acl/user/application/contract"
	appmock "goedabook/test/mock"
)

func TestSignupHandler(t *testing.T) {
	t.Run("ensure that create a user", func(tr *testing.T) {
		type fields struct {
			validators []contract.Validator
		}

		type args struct {
			ctx context.Context
			cmd contract.Command
		}

		type exp struct {
			want contract.HandlerResult
			err  contract.HandlerErr
		}

		tests := []struct {
			name   string
			fields *fields
			args   *args
			exp    *exp
		}{
			{
				name:   "should result a handler error when required field is not provided",
				fields: &fields{
					validators: []contract.Validator{
						func(t *testing.T) contract.Validator {
							v := appmock.NewValidator(t)
							v.On("Validate", mock.Anything).Return(fmt.Errorf(RequiredFieldErrMsg, "username"))
							return v
						}(tr),
					},
				},
				args: &args{
					ctx: context.Background(),
					cmd: func(t *testing.T) contract.Command {
						cmd := appmock.NewCommand(t)
						return cmd
					}(tr),
				},
				exp: &exp{
					want: nil,
					err:  fmt.Errorf(RequiredFieldErrMsg, "username"),
				},
			},
			{
				name:   "should receive a handler result when all fields are provided",
				fields: &fields{
					validators: []contract.Validator{
						func(t *testing.T) contract.Validator {
							v := appmock.NewValidator(t)
							v.On("Validate", mock.Anything).Return(nil)
							return v
						}(tr),
					},
				},
				args: &args{
					ctx: context.Background(),
					cmd: func(t *testing.T) contract.Command {
						cmd := appmock.NewCommand(t)
						return cmd
					}(tr),
				},
				exp: &exp{
					want: nil,
					err:  nil,
				},
			},
		}

		for _, tt := range tests {
			tr.Run(tt.name, func(tr *testing.T) {
				sut := NewSignupHandler(WithValidators(tt.fields.validators...))

				got, err := sut.Handle(tt.args.ctx, tt.args.cmd)

				if (err != nil) && (tt.exp.err != nil) && (err.Error() != tt.exp.err.Error()) {
					tr.Errorf("expected error: %v, got: %v", tt.exp.err, err)
				}

				if got != tt.exp.want{
					tr.Errorf("expected result: %v, got: %v", tt.exp.want, got)
				}
			})
		}
	})
}
