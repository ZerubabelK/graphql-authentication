package graphqlClient

import (
	graphql "github.com/hasura/go-graphql-client"
)

var INSERT_ONE_MUTATION struct {
	InsertUserOne struct {
		ID graphql.ID
	} `graphql:"insert_user_one(object: $object)"`
}

var GetUserByUsernameQuery struct {
	User []struct {
		ID            graphql.ID      `graphql:"id"`
		Password      graphql.String  `graphql:"password"`
		EmailVerified graphql.Boolean `graphql:"email_verified"`
		Role          graphql.String  `graphql:"role"`
	} `graphql:"user(where: { username: { _eq: $username } })"`
}

var GetUserByIDQuery struct {
	User []struct {
		ID       graphql.ID
		Email    graphql.String
		Username graphql.String
	} `graphql:"user(where: { id: { _eq: $id } })"`
}

var UpdateProfileImageMutation struct {
	UpdateUserByID struct {
		ID           graphql.ID     `graphql:"id"`
		Email        graphql.String `graphql:"email"`
		FirstName    graphql.String `graphql:"first_name"`
		LastName     graphql.String `graphql:"last_name"`
		ProfileImage graphql.String `graphql:"profile_image"`
	} `graphql:"update_user_by_pk(pk_columns: { id: $id }, _set: { profile_image: $profile_image })"`
}

var UpdateUserByIDMutation struct {
	UpdateUserByID struct {
		ID graphql.ID
	} `graphql:"update_user_by_pk(pk_columns: { id: $id }, _set: $object)"`
}

var CREATE_NOTIFICATION_MUTATION struct {
	InsertNotificationOne struct {
		ID graphql.ID
	} `graphql:"insert_notification_one(object: $object)"`
}

var USER_BY_PK struct {
	User struct {
		ID        graphql.ID     `graphql:"id"`
		Email     graphql.String `graphql:"email"`
		FirstName graphql.String `graphql:"first_name"`
		LastName  graphql.String `graphql:"last_name"`
		Username  graphql.String `graphql:"username"`
	} `graphql:"user_by_pk(id: $id)"`
}

var GetRecipeCreatorFromLikeIdQuery struct {
	Like struct {
		User struct {
			Username graphql.String `graphql:"username"`
		} `graphql:"user"`
		Recipe struct {
			UserID graphql.ID `graphql:"user_id"`
		} `graphql:"recipe"`
	} `graphql:"like_by_pk(id: $id)"`
}

var GetRecipeCreator struct {
	Recipe struct {
		User struct {
			UserID   graphql.ID     `graphql:"id"`
			Username graphql.String `graphql:"username"`
		} `graphql:"user"`
	} `graphql:"recipe_by_pk(id: $id)"`
}

var GetUserByRole struct {
	User []struct {
		ID graphql.ID `graphql:"id"`
	} `graphql:"user(where: { role: { _eq: $role } })"`
}
