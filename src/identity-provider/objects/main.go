package objects

import (
	_ "encoding/json"
)

type User struct {
	Id        int    `json:"id" gorm:"primary_key; index"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" gorm:"not null; unique"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"not null; unique"`
	UserType  string `json:"user_type" gorm:"not null" sql:"DEFAULT: 'user'"`
}

func (User) TableName() string {
	return "user"
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
