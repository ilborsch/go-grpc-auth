package tkn

import (
	"github.com/golang-jwt/jwt/v5"
	"go-grpc-auth/internal/models"
	"time"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["app_id"] = app.ID
	claims["expiration"] = time.Now().Add(duration).Unix()

	tokenSigned, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}
