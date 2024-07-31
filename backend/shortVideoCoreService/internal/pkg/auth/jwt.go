package auth

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId int64, config *conf.Common) (string, error) {
	hours := time.Duration(config.Auth.Expire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * hours).Unix(),
	})
	signedToken, err := token.SignedString([]byte(config.Auth.Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
