package structs

import (
	uuid "github.com/google/uuid"
)

type SportmenSport struct {
	Sport string `json:"sport" validate:"oneof=basketball cycling"`
}

type Sportmen struct {
	UserId           uuid.UUID       `json:"user"`
	Name             string          `json:"name" validate:"required"`
	LastName         string          `json:"last_name" validate:"required"`
	Age              int             `json:"age" validate:"required"`
	Weight           int             `json:"weight" validate:"required"`
	Height           int             `json:"height" validate:"required"`
	CountryBirth     string          `json:"country_birth" validate:"required"`
	CityBirth        string          `json:"city_birth" validate:"required"`
	CountryResidence string          `json:"country_residence" validate:"required"`
	CityResidence    string          `json:"city_residence" validate:"required"`
	LengthResidence  int             `json:"length_residence" validate:"required"`
	Sport            []SportmenSport `validate:"dive"`
}
