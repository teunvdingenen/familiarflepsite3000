package resolvers

import (
	"github.com/graph-gophers/graphql-go"
	log "github.com/sirupsen/logrus"
	"github.com/teunvdingenen/familiarflepsite3000/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpInput struct {
	Firstname string
	Lastname  string
}

func (r *Resolver) SignUp(args struct{ SignUpInput *SignUpInput }) (graphql.ID, error) {
	objectID, err := r.Datasource.GetUserRepository().SaveUser(data.User{
		ID:        primitive.NewObjectID(),
		Firstname: args.SignUpInput.Firstname,
		Lastname:  args.SignUpInput.Lastname,
	})
	if err != nil {
		log.Error(err)
		return graphql.ID(""), err
	}
	return graphql.ID(objectID.Hex()), nil
}
