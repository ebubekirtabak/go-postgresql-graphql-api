package types

import (
	"github.com/graphql-go/graphql"
	"password-manager-server/common"
	"go-postgresql-graphql-api/database"
	"fmt"
)


// Role type definition.
type Comment struct {
	ID   	 int    `db:"id" json:"id"`
	PostId   int    `db:"postId" json:"postId"`
	Name 	 string `db:"name" json:"name"`
	Email 	 string `db:"email" json:"email"`
	Body     string `db:"body" json:"body"`
}

func CreateCommentType () *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"postid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func fetchCommentsByPostID(id int) ([]Comment, error) {
	var response []Comment
	db := database.OpenDatabase()
	defer db.Close()
	rows, err := db.Query(fmt.Sprintf("SELECT id, postid, name, email, body FROM comment WHERE postid='%d'", id))
	common.CheckErr(err)
	isNext := true

	for rows.Next() {
		commentData := Comment{}
		isNext = false
		rows.Scan(&commentData.ID, &commentData.PostId, &commentData.Name, &commentData.Email, &commentData.Body)
		response = append(response, commentData)
	}

	if isNext {
		return nil, nil
	} else {
		return response, nil
	}
}


