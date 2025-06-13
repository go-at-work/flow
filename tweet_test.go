package flow

import (
	"testing"

	"github.com/arisromil/flow/faker"
	"github.com/stretchr/testify/require"
)

func TestCreateTweetRequest_Sanitize(t *testing.T) {
	t.Parallel()

	input := CreateTweetRequest{
		Body: " Hello  ",
	}
	want := CreateTweetRequest{
		Body: "Hello",
	}

	input.Sanitize()

	require.Equal(t, want.Body, input.Body, "Sanitize should trim whitespace from the body")
}

func TestCreateTweetRequest_Validate(t *testing.T) {
	t.Parallel()

	ErrValidation := 0
	testCases := []struct {
		name  string
		input CreateTweetRequest
		err   error
	}{
		{
			name: "valid input",
			input: CreateTweetRequest{
				Body: "This is a valid tweet body.",
			},
			err: nil,
		},
		{
			name: "empty body",
			input: CreateTweetRequest{
				Body: "",
			},
			err: ErrValidation,
		},
		{
			name: "body too long",
			input: CreateTweetRequest{
				Body: faker.RandomString(300),
			},
			err: ErrValidation,
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
