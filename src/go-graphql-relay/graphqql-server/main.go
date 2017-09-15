package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
	_ "fmt"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams,) ( interface{}, error) {
				return "Hello world!", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})
	
	// serve HTTP
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}