package main

import (
	"log"
	"net/http"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"go-postgresql-graphql-api/types"
	"go-postgresql-graphql-api/database"
)



func main() {
	database.TestDatabase()
	database.OpenDatabase()
	schemaPost, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQueryType(
					createPostType(
					createCommentType(),
				),
			),

		),
	})

	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}

	handlerPost := gqlhandler.New(&gqlhandler.Config {
		Schema: &schemaPost,
	})

	http.Handle("/post", handlerPost)
	log.Println("Server started at http://localhost:3001/post")
	log.Fatal(http.ListenAndServe(":3001", nil))
}


func createQueryType(postType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: postType,
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				v, _ := id.(int)
				log.Printf("fetching post with id: %d", v)
				return types.FetchPostByiD(v)
			},
		},
	}}
}

func createPostType(commentType *graphql.Object) *graphql.Object {
	return types.CreatePostType(commentType)
}

func createCommentType() *graphql.Object {
	return types.CreateCommentType()
}




