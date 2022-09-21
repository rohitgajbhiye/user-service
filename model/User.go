package model

type Users map[string]User

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}
