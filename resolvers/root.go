package resolvers

import (
	"github.com/teunvdingenen/familiarflepsite3000/data"
)

type Resolver struct {
	Datasource *data.MongoDatastore
}

//All data queries start here
func (r *Resolver) Viewer() *UserResolver {
	//TODO get logged in id

	user, err := r.Datasource.GetUserRepository().GetOneUser()
	if err != nil {
		return nil
	}
	return &UserResolver{u: user}
}
