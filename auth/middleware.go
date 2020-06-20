package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net"
	"net/http"
	"strings"
)

type UserAuth struct {
	UserID    []byte
	Roles     []string
	IPAddress string
	Token     string
}

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
type UserClaims struct {
	UserId []byte `json:"user_id"`
	jwt.StandardClaims
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := TokenFromHttpRequest(r)

		userId := UserIDFromToken(token)

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		userAuth := UserAuth{
			UserID:    userId,
			IPAddress: ip,
			Token:     token,
		}

		ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func TokenFromHttpRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	var tokenString string
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}
	return tokenString
}
func UserIDFromToken(tokenString string) []byte {
	token, err := JwtDecode(tokenString)
	if err != nil {
		return nil
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims == nil {
			return nil
		}
		return claims.UserId
	} else {
		return nil
	}
}
func ForContext(ctx context.Context) *UserAuth {
	raw := ctx.Value(userCtxKey)
	if raw == nil {
		return nil
	}
	return raw.(*UserAuth)
}
func GetAuthFromContext(ctx context.Context) *UserAuth {
	return ForContext(ctx)
}
