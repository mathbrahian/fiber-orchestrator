package structs

import (
	uuid "github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email" validate:"required"`
	Password string    `json:"password" validate:"required"`
}

type UserDelete struct {
	ID uuid.UUID `json:"id"`
}

type UserRespose struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
