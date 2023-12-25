package handlers

import (
	"context"
	graphqlClient "graphql/internal/graphql"
	"graphql/internal/types"
	"graphql/internal/util"
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
				"message": err,
			})
			return
		}

		client := graphqlClient.Client()

		variables := map[string]interface{}{
			"username": input.Input.Credential.Username,
		}
		err = client.Query(context.Background(), &graphqlClient.GetUserByUsernameQuery, variables)

		if err != nil || len(graphqlClient.GetUserByUsernameQuery.User) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		if !graphqlClient.GetUserByUsernameQuery.User[0].EmailVerified {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Email not verified",
			})
			return
		}

		valid := util.CheckPasswordHash(input.Input.Credential.Password, string(graphqlClient.GetUserByUsernameQuery.User[0].Password))

		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid username or password"})
			return
		}

		token, err := jwt_modules.GenerateToken(string(graphqlClient.GetUserByUsernameQuery.User[0].ID), string(graphqlClient.GetUserByUsernameQuery.User[0].Role))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
			"success":     true,
			"error":       nil,
		})

	}
}

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input types.RegisterInput
		err := c.BindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		client := graphqlClient.Client()

		type user_insert_input struct {
			FirstName string `graphql:"first_name" json:"first_name"`
			LastName  string `graphql:"last_name" json:"last_name"`
			Email     string `graphql:"email" json:"email"`
			Username  string `graphql:"username" json:"username"`
			Password  string `graphql:"password" json:"password"`
		}

		hashedPassword, err := util.HashPassword(input.Input.User.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		variables := map[string]interface{}{
			"object": user_insert_input{
				FirstName: input.Input.User.FirstName,
				LastName:  input.Input.User.LastName,
				Email:     input.Input.User.Email,
				Username:  input.Input.User.Username,
				Password:  hashedPassword,
			},
		}

		err = client.Mutate(context.Background(), &graphqlClient.INSERT_ONE_MUTATION, variables)

		log.Println(graphqlClient.INSERT_ONE_MUTATION.InsertUserOne)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "username or email already exists",
			})
			return
		}

		// token, err := jwt_modules.GenerateToken(string(graphqlClient.INSERT_ONE_MUTATION.InsertUserOne.ID))

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"message": "Internal Server Error",
		// 	})
		// 	return
		// }

		verificationToken, er := jwt_modules.GenerateVerificationToken(string(graphqlClient.INSERT_ONE_MUTATION.InsertUserOne.ID))

		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		emailBody, err := util.GetEmailBody(verificationToken)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		err = util.SendEmail(emailBody, input.Input.User.Email, "Email Verification | Tastebite")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		log.Println("email", err)

		c.JSON(http.StatusOK, gin.H{
			"user_id": string(graphqlClient.INSERT_ONE_MUTATION.InsertUserOne.ID),
			"success": true,
			"error":   nil,
		})
	}
}

func VerifyEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input types.VerificationInput
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		claim, ok := jwt_modules.VerifyVerificationToken(input.Input.Arg.Token)
		if ok != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		client := graphqlClient.Client()

		type user_set_input struct {
			EmailVerified bool `graphql:"email_verified" json:"email_verified"`
		}

		type uuid string

		variables := map[string]interface{}{
			"id": uuid(claim.UserID),
			"object": user_set_input{
				EmailVerified: true,
			},
		}

		err = client.Mutate(context.Background(), &graphqlClient.UpdateUserByIDMutation, variables)
		log.Println(err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})

	}
}

func ResendVerificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input types.ResendVerificationInput
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		client := graphqlClient.Client()

		type uuid string
		variables := map[string]interface{}{
			"id": uuid(input.Input.Arg.UserID),
		}

		err = client.Query(context.Background(), &graphqlClient.GetUserByIDQuery, variables)

		if err != nil || len(graphqlClient.GetUserByIDQuery.User) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		verificationToken, er := jwt_modules.GenerateVerificationToken(string(graphqlClient.GetUserByIDQuery.User[0].ID))

		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		emailBody, err := util.GetEmailBody(verificationToken)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		err = util.SendEmail(emailBody, string(graphqlClient.GetUserByIDQuery.User[0].Email), "Email Verification | Tastebite")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
