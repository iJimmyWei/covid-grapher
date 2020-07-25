package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"

	"github.com/ijimmywei/covid-grapher/cors"
	"github.com/ijimmywei/covid-grapher/db"
	"github.com/ijimmywei/covid-grapher/graph/generated"
	"github.com/ijimmywei/covid-grapher/graph/resolver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	http.Handle("/query", gqlHandler(db.New(client)))
	http.Handle("/", playground.Handler("Covid", "/query"))
	err = http.ListenAndServe(":8085", nil)
	log.Println(err)
}

func gqlHandler(db db.DB) http.HandlerFunc {
	config := generated.Config{
		Resolvers: &resolver.Resolver{DB: db},
	}
	gh := handler.GraphQL(generated.NewExecutableSchema(config))
	if os.Getenv("profile") != "prod" {
		gh = cors.Disable(gh)
	}
	return gh
}

func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
