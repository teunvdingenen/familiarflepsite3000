package main

import (
	"github.com/friendsofgo/graphiql"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"
	"github.com/teunvdingenen/familiarflepsite3000/auth"
	"github.com/teunvdingenen/familiarflepsite3000/config"
	"github.com/teunvdingenen/familiarflepsite3000/data"
	"github.com/teunvdingenen/familiarflepsite3000/resolvers"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	s, err := getSchema("./schema.graphqls")
	if err != nil {
		panic(err)
	}

	logger := log.New()
	datastore := data.NewDatastore(&config.Config, logger)
	rootResolver := resolvers.Resolver{datastore}
	schema := graphql.MustParseSchema(s, &rootResolver)
	http.Handle("/query", &relay.Handler{Schema: schema})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()

	mux.Handle("/graphql", auth.Middleware(&relay.Handler{Schema: schema}))
	mux.Handle("/graphiql", graphiqlHandler)

	logger.WithFields(log.Fields{"time": time.Now()}).Info("starting server")
	logger.Fatal(http.ListenAndServe("localhost:3000", logged(mux)))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// logging middleware
func logged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"path":    r.RequestURI,
			"IP":      r.RemoteAddr,
			"elapsed": time.Now().UTC().Sub(start),
		}).Info()
	})
}
