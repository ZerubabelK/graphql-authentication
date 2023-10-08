package jwt_modules

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyJWTClaims struct {
	Payload struct {
		AllowedRoles []string `json:"x-hasura-allowed-roles"`
		DefaultRole  string   `json:"x-hasura-default-role"`
		UserID       string   `json:"x-hasura-user-id"`
	} `json:"https://hasura.io/jwt/claims"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := MyJWTClaims{
		Payload: struct {
			AllowedRoles []string `json:"x-hasura-allowed-roles"`
			DefaultRole  string   `json:"x-hasura-default-role"`
			UserID       string   `json:"x-hasura-user-id"`
		}{
			AllowedRoles: []string{"user"},
			DefaultRole:  "user",
			UserID:       userId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "https://hasura.io/jwt/claims",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func VerifyToken(token string) (*MyJWTClaims, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := &MyJWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return claims, nil
}
