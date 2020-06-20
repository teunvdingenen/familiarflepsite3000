package resolvers

import (
	"context"
	"errors"
	"github.com/teunvdingenen/familiarflepsite3000/auth"
	"github.com/teunvdingenen/familiarflepsite3000/config"
	"time"
)

type LoginInput struct {
	Email    string
	Password string
}



func (r *Resolver) Login(args struct{ LoginInput *LoginInput }) (*TokenResolver, error) {
	user, err := r.Datasource.GetUserRepository().GetUserByEmail(context.TODO(), args.LoginInput.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}
	if !auth.ComparePassword(args.LoginInput.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	expiredAt := time.Now().Add(config.Config.Auth.ValidDuration).Unix()
	obj := Token {
		Token:	auth.JwtCreate([]byte(user.ID.Hex()), expiredAt),
		ExpiredAt: int32(expiredAt),
	}
	return &TokenResolver{&obj}, nil
}