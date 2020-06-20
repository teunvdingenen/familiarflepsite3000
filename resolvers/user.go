package resolvers

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/teunvdingenen/familiarflepsite3000/data"
)

type UserResolver struct {
	u *data.User
}

func (u *UserResolver) ID() graphql.ID {
	return graphql.ID(u.u.ID.Hex())
}

func (u *UserResolver) Firstname() *string {
	firstname := &u.u.Firstname
	if *firstname == "" {
		return nil
	}
	return firstname
}

func (u *UserResolver) Lastname() *string {
	lastname := &u.u.Lastname
	if *lastname == "" {
		return nil
	}
	return lastname
}

func (u *UserResolver) Email() *string {
	email := &u.u.Email
	if *email == "" {
		return nil
	}
	return email
}
