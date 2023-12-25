package handlers

import (
	"context"
	graphqlClient "graphql/internal/graphql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LikeNotificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload struct {
			Event struct {
				Op   string `json:"op"`
				Data struct {
					New struct {
						UserID    string `json:"user_id"`
						RecipeId  string `json:"recipe_id"`
						CreatedAt string `json:"created_at"`
						ID        string `json:"id"`
					} `json:"new"`
				} `json:"data"`
			} `json:"event"`
		}

		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		client := graphqlClient.Client()

		type uuid string

		variables := map[string]interface{}{
			"id": uuid(payload.Event.Data.New.ID),
		}

		err = client.Query(context.Background(), &graphqlClient.GetRecipeCreatorFromLikeIdQuery, variables)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		log.Println(graphqlClient.GetRecipeCreatorFromLikeIdQuery.Like.Recipe.UserID)

		// err = client.Query(context.Background(), &graphqlClient.USER_BY_PK, variables)
		// if err != nil {
		// 	c.JSON(http.StatusNotFound, gin.H{
		// 		"message": err,
		// 	})
		// 	return
		// }

		type notification_insert_input struct {
			UserId      uuid   `graphql:"user_id" json:"user_id"`
			Message     string `graphql:"message" json:"message"`
			InitiatorId uuid   `graphql:"initiator_id" json:"initiator_id"`
		}

		variables = map[string]interface{}{
			"object": notification_insert_input{
				UserId:      uuid(graphqlClient.GetRecipeCreatorFromLikeIdQuery.Like.Recipe.UserID),
				Message:     "user @" + string(graphqlClient.GetRecipeCreatorFromLikeIdQuery.Like.User.Username) + " liked your recipe",
				InitiatorId: uuid(payload.Event.Data.New.UserID),
			},
		}

		err = client.Mutate(context.Background(), &graphqlClient.CREATE_NOTIFICATION_MUTATION, variables)

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

func CommentNotificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AdminRecipeNotificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}
		recipe_id := body.(map[string]interface{})["event"].(map[string]interface{})["data"].(map[string]interface{})["new"].(map[string]interface{})["id"].(string)
		// var payload struct {
		// 	Event struct {
		// 		Op   string `json:"op"`
		// 		Data struct {
		// 			New struct {
		// 				UserID      string `json:"user_id"`
		// 				Description string `json:"description"`
		// 				TItle       string `json:"title"`
		// 				PrepTime    string `json:"prep_time"`
		// 				CreatedAt   string `json:"created_at"`
		// 				UpdatedAt string `json:"updated_at"`
		// 				RecipeCategoryId string `json:"recipe_category_id"`
		// 				IsVerified  bool   `json:"is_verified"`
		// 				ID          string `json:"id"`
		// 			} `json:"new"`
		// 		} `json:"data"`
		// 	} `json:"event"`
		// }

		// err := c.ShouldBindJSON(&payload)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"message": err,
		// 	})
		// 	return
		// }

		client := graphqlClient.Client()

		type uuid string

		variables := map[string]interface{}{
			"id": uuid(recipe_id),
		}

		err := client.Query(context.Background(), &graphqlClient.GetRecipeCreator, variables)

		log.Println(err)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		log.Println(graphqlClient.GetRecipeCreator.Recipe.User)

		err = client.Query(context.Background(), &graphqlClient.GetUserByRole, map[string]interface{}{"role": "admin"})
		if err != nil || len(graphqlClient.GetUserByRole.User) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}
		log.Println(graphqlClient.GetUserByRole.User)

		type notification_insert_input struct {
			UserId  uuid   `graphql:"user_id" json:"user_id"`
			Message string `graphql:"message" json:"message"`
		}

		variables = map[string]interface{}{
			"object": notification_insert_input{
				UserId:  uuid(graphqlClient.GetUserByRole.User[0].ID),
				Message: "New Recipe Verification Required for @" + string(graphqlClient.GetRecipeCreator.Recipe.User.Username) + "'s recipe",
			},
		}

		err = client.Mutate(context.Background(), &graphqlClient.CREATE_NOTIFICATION_MUTATION, variables)
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
