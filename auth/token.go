package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/teunvdingenen/familiarflepsite3000/config"
)

func JwtDecode(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Config.Auth.SigningKey, nil
	})
}

func JwtCreate(userId []byte, expiredAt int64) string {
	claims := UserClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer: config.Config.Auth.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(config.Config.Auth.SigningKey)
	return ss
}