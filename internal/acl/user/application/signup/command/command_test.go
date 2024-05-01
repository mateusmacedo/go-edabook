package command

import (
	"fmt"
	"reflect"
	"testing"

	"goedabook/internal/acl/user/domain/valueobj"
)

func TestSignupCommmand(t *testing.T) {
	t.Run("should create a new signup command", func(tr *testing.T) {
		type fields struct {
		}

		type args struct {
			options []SignupCommandOptions
		}

		type exp struct {
			want *SignupCommand
			err  error
		}

		tests := []struct {
			name   string
			fields *fields
			args   *args
			exp    *exp
		}{
			{
				name:   "should create a new signup command",
				fields: &fields{},
				args: &args{
					options: []SignupCommandOptions{
						WithUserName(valueobj.Username("username")),
					},
				},
				exp: &exp{
					want: &SignupCommand{
						Username: valueobj.Username("username"),
					},
					err: nil,
				},
			},
			{
				name:   "should return an error when option func returns an error",
				fields: &fields{},
				args: &args{
					options: []SignupCommandOptions{
						func(cmd *SignupCommand) error {
							return fmt.Errorf("error")
						},
					},
				},
				exp: &exp{
					want: &SignupCommand{
						Username: valueobj.Username("username"),
					},
					err:  fmt.Errorf("error"),
				},
			},
		}

		for _, tt := range tests {
			tr.Run(tt.name, func(tr *testing.T) {
				got, err := NewSignUpCommand(tt.args.options...)

				if err != nil {
					if tt.exp.err == nil {
						tr.Errorf("NewSignUpCommand() error = %v, wantErr %v", err, tt.exp.err)
						return
					}
				}

				if got != nil && !reflect.DeepEqual(got, tt.exp.want) {
					tr.Errorf("NewSignUpCommand() = %v, want %v", got, tt.exp.want)
				}
			})
		}
	})
}
