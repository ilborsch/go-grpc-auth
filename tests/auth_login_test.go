package tests

import (
	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"github.com/brianvoe/gofakeit"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-grpc-auth/tests/suite"
	"testing"
	"time"
)

type LoginTestCase struct {
	name        string
	email       string
	password    string
	appID       int32
	expectedErr string
}

func TestLogin_HappyLogin(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 12)

	responseReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	require.NoErrorf(t, err, "register error")
	assert.NotEmptyf(t, responseReg.GetUserId(), "register uid is empty")

	responseLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    appID,
	})
	require.NoErrorf(t, err, "login error")

	token := responseLogin.GetToken()
	require.NotEmptyf(t, token, "login token is empty")

	tokenDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoErrorf(t, err, "error parsing login token")

	claims, ok := tokenDecoded.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, int64(claims["uid"].(float64)), responseReg.GetUserId())
	assert.Equal(t, claims["email"], email)
	assert.Equal(t, int(claims["app_id"].(float64)), appID)

	expectedExpiration := time.Now().Add(st.Cfg.GRPC.Timeout).Unix()
	// TODO: this test fails
	// test if token expiration is correct (1 second delta is acceptable)
	assert.InDelta(t, expectedExpiration, claims["expiration"].(float64), 1)
}

func TestLogin_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []LoginTestCase{
		{
			name:        "Login with empty email",
			email:       "",
			password:    gofakeit.Password(true, true, true, true, false, 12),
			appID:       appID,
			expectedErr: "email is required",
		},
		{
			name:        "Login with empty password",
			email:       gofakeit.Email(),
			password:    "",
			appID:       appID,
			expectedErr: "password is required",
		},
		{
			name:        "Login with empty email and password",
			email:       "",
			password:    "",
			appID:       appID,
			expectedErr: "email is required",
		},
		{
			name:        "Login with non existing credentials",
			email:       gofakeit.Email(),
			password:    gofakeit.Password(true, true, true, true, false, 12),
			appID:       appID,
			expectedErr: "internal error",
		},
		{
			name:        "Login without appID",
			email:       gofakeit.Email(),
			password:    gofakeit.Password(true, true, true, true, false, 12),
			appID:       emptyAppID,
			expectedErr: "app_id is required",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
				Email:    test.email,
				Password: test.password,
				AppId:    test.appID,
			})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.expectedErr)
		})
	}

}
