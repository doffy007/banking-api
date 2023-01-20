package dto

import "time"

type UserDTO struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
}

type UserDTOForLogin struct {
	ID int `json:"id"`
}

type UserWithTokenDTO struct {
	User  UserDTOForLogin `json:"user"`
	Token string          `json:"token"`
}
