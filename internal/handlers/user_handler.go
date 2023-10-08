package handlers

import (
	"context"
	graphqlClient "graphql/internal/graphql"
	"graphql/internal/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the JWT
		userID := c.MustGet("userID").(string)

		// Get the user from the database
		client := graphqlClient.Client()
		variables := map[string]interface{}{
			"id": userID,
		}
		err := client.Query(context.Background(), &graphqlClient.GetUserByIDQuery, variables)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "User not found",
			})
			return
		}

		// Return the user
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"user":    graphqlClient.GetUserByIDQuery.User[0],
		})
	}
}

type UploadInput struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Arg struct {
			Base64 string `json:"base64"`
			Name   string `json:"name"`
		} `json:"arg"`
	} `json:"input"`
}

func UpdateProfileImageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the file from the request
		var profile_image UploadInput
		err := c.BindJSON(&profile_image)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"success":           false,
				"error":             "Invalid file",
				"profile_image_url": nil,
			})
			return
		}

		// Upload the file to Cloudinary
		profile_image_url, err := util.UploadFile(profile_image.Input.Arg.Base64, profile_image.Input.Arg.Name)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":           false,
				"error":             "Error uploading file",
				"profile_image_url": nil,
			})
			return
		}

		// // Return the user
		log.Println(graphqlClient.UpdateProfileImageMutation.UpdateUserByID)
		c.JSON(http.StatusOK, gin.H{
			"success":           true,
			"error":             nil,
			"profile_image_url": profile_image_url,
		})

	}
}
