package handlers

import (
	"context"
	graphqlClient "graphql/internal/graphql"
	"graphql/internal/types"
	jwt_modules "graphql/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input types.LoginInput
		err := c.BindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Bad Request",
			})
			return
		}

		client := graphqlClient.Client()

		variables := map[string]interface{}{
			"email": input.Input.Credential.Email,
		}
		err = client.Query(context.Background(), &graphqlClient.GetUserByEmailQuery, variables)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Internal Server Error",
			})
			return
		}

		if len(graphqlClient.GetUserByEmailQuery.User) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Email not found",
			})
			return
		}

		token, err := jwt_modules.GenerateToken(string(graphqlClient.GetUserByEmailQuery.User[0].ID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
			"success":     true,
			"error":       "",
		})

	}
}

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input types.RegisterInput
		err := c.BindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Bad Request",
			})
			return
		}

		client := graphqlClient.Client()

		type user_insert_input struct {
			Email     string `graphql:"email" json:"email"`
			FirstName string `graphql:"first_name" json:"first_name"`
			LastName  string `graphql:"last_name" json:"last_name"`
			Password  string `graphql:"password" json:"password"`
		}

		variables := map[string]interface{}{
			"object": user_insert_input{
				Email:     input.Input.User.Email,
				FirstName: input.Input.User.FirstName,
				LastName:  input.Input.User.LastName,
				Password:  input.Input.User.Password,
			},
		}

		err = client.Mutate(context.Background(), &graphqlClient.INSERT_ONE_MUTATION, variables)

		log.Println(graphqlClient.INSERT_ONE_MUTATION.InsertUserOne)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Email already exists",
			})
			return
		}

		token, err := jwt_modules.GenerateToken(string(graphqlClient.INSERT_ONE_MUTATION.InsertUserOne.ID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"accessToken": "",
				"success":     false,
				"error":       "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
			"success":     true,
			"error":       "",
		})
	}
}
