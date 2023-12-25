package types

import (
	"github.com/hasura/go-graphql-client"
)

type User struct {
	ID            graphql.ID     `graphql:"id" json:"id"`
	FirstName     string         `graphql:"first_name" json:"first_name"`
	LastName      string         `graphql:"last_name" json:"last_name"`
	Email         string         `graphql:"email" json:"email"`
	Username      string         `graphql:"username" json:"username"`
	CreatedAt     graphql.String `graphql:"created_at" json:"created_at,omitempty"`
	UpdatedAt     graphql.String `graphql:"updated_at" json:"updated_at,omitempty"`
	ProfileImage  string         `graphql:"profile_image" json:"profile_image,omitempty"`
	EmailVerified bool           `graphql:"email_verified" json:"email_verified,omitempty"`
}

type LoginInput struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Credential struct {
			Username string `json:"username"`
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
			Username  string `json:"username"`
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

type VerificationInput struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Arg struct {
			Token string `json:"token"`
		} `json:"arg"`
	} `json:"input"`
}

type ResendVerificationInput struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Arg struct {
			UserID string `json:"user_id"`
		} `json:"arg"`
	} `json:"input"`
}
