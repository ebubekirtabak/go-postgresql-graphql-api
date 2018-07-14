package types

import (
	"github.com/graphql-go/graphql"
	"fmt"
	"log"
	"go-postgresql-graphql-api/database"
)


// Role type definition.
type Post struct {
	ID   	 int    `db:"id" json:"id"`
	UserId   int    `db:"userid" json:"userId"`
	Title 	 string `db:"title" json:"title"`
	Body     string `db:"body" json:"body"`
}

func CreatePostType(commentType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.Int},
			"userId": &graphql.Field{Type: graphql.Int},
			"title":  &graphql.Field{Type: graphql.String},
			"body":   &graphql.Field{Type: graphql.String},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post, _ := p.Source.(Post)
					return fetchCommentsByPostID(post.ID)
				},
			},
		},
	})
}

func FetchPostByiD(id int) (Post, error) {
	postData := Post{}
	db := database.OpenDatabase()
	defer db.Close()
	err := db.QueryRow(fmt.Sprintf("select id, userid, title, body from post where id='%d';",id)).Scan(&postData.ID, &postData.UserId, &postData.Title, &postData.Body)

	if err != nil {
		log.Fatal(err)
	}

	return postData,err
}






