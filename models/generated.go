// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
}

type AuthToken struct {
	AccessToken string `json:"accessToken"`
	ExpireAt    string `json:"expireAt"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
