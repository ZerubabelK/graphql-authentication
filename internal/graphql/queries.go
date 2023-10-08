package graphqlClient

import (
	graphql "github.com/hasura/go-graphql-client"
)

var INSERT_ONE_MUTATION struct {
	InsertUserOne struct {
		ID graphql.ID
	} `graphql:"insert_user_one(object: $object)"`
}

var GetUserByEmailQuery struct {
	User []struct {
		ID       graphql.ID
		Password graphql.String
	} `graphql:"user(where: { email: { _eq: $email } })"`
}

var GetUserByIDQuery struct {
	User []struct {
		ID        graphql.ID
		Email     graphql.String
		FirstName graphql.String
		LastName  graphql.String
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
	} `graphql:"update_user_by_pk(pk_columns: { id: $id }, _set: { email: $email, first_name: $first_name, last_name: $last_name })"`
}
