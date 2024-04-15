package tests

import (
	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-grpc-auth/tests/suite"
	"testing"
)

const UID = 1

type IsAdminTestCase struct {
	name        string
	userID      int64
	expectedErr string
}

func TestIsAdmin_HappyAuthorization(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 12)

	responseReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, responseReg.GetUserId())

	responseAuth, err := st.AuthClient.IsAdmin(ctx, &ssov1.IsAdminRequest{
		UserId: responseReg.GetUserId(),
	})
	require.NoError(t, err)
	assert.Equalf(t, false, responseAuth.IsAdmin, "user is not admin")
}

func TestIsAdmin_FailCases(t *testing.T) {
	ctx, st := suite.New(t)
	tests := []IsAdminTestCase{
		{
			name:        "Non-existing UID",
			userID:      int64(1) << 62,
			expectedErr: "internal error",
		},
		{
			name:        "Negative UID",
			userID:      -1,
			expectedErr: "user id cannot be negative",
		},
		{
			name:        "Empty UID",
			userID:      0,
			expectedErr: "user id is required",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := st.AuthClient.IsAdmin(ctx, &ssov1.IsAdminRequest{
				UserId: test.userID,
			})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.expectedErr)
		})
	}
}
