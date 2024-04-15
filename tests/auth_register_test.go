package tests

import (
	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-grpc-auth/tests/suite"
	"testing"
)

const (
	appID      = 1
	emptyAppID = 0
	appSecret  = "testSecret"
)

type RegisterTestCase struct {
	name        string
	email       string
	password    string
	expectedErr string
}

func TestRegister_DuplicatedRegister(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 12)

	registerRes, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	require.NoErrorf(t, err, "register error")
	require.NotEmptyf(t, registerRes.GetUserId(), "register uid is empty")

	registerRes, err = st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	require.Error(t, err)
	require.Empty(t, registerRes.GetUserId())
}

func TestRegister_FailCases(t *testing.T) {
	ctx, st := suite.New(t)
	tests := []RegisterTestCase{
		{
			name:        "Register with empty email",
			email:       "",
			password:    gofakeit.Password(true, true, true, true, false, 12),
			expectedErr: "email is required",
		},
		{
			name:        "Register with empty password",
			email:       gofakeit.Email(),
			password:    "",
			expectedErr: "password is required",
		},
		{
			name:        "Register with both empty",
			email:       "",
			password:    "",
			expectedErr: "email is required",
		},
		{
			name:        "Password is too small",
			email:       gofakeit.Email(),
			password:    gofakeit.Password(true, true, true, true, false, 7),
			expectedErr: "password is too small",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    test.email,
				Password: test.password,
			})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.expectedErr)
		})
	}
}
