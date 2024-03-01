package structs

import (
	uuid "github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email" validate:"required"`
	Password string    `json:"password" validate:"required"`
}
