package flow

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid input",
			input: RegisterInput{
				Username:        "testuser",
				Email:           "testuser@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "too short username",
			input: RegisterInput{
				Username:        "t",
				Email:           "testuser@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: NewValidationError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
