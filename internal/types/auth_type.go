package types

import (
	"github.com/hasura/go-graphql-client"
)

type User struct {
	ID        graphql.ID `graphql:"id" json:"id"`
	Email     string     `graphql:"email" json:"email"`
	FirstName string     `graphql:"first_name" json:"first_name"`
	LastName  string     `graphql:"last_name" json:"last_name"`
}

type LoginInput struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Credential struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"credential"`
	} `json:"input"`
}

type RegisterInput struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		User struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		} `json:"user"`
	} `json:"input"`
}

type AuthResponse struct {
	Token   string `json:"accessToken"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
