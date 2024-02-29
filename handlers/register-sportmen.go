package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

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

func RegisterSportmen(c *fiber.Ctx) error {

	NewSportmen := new(Sportmen)

	if err := c.BodyParser(NewSportmen); err != nil {
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(NewSportmen)

	err2 := validate.Struct(NewSportmen.Sport)

	if err1 != nil && err2 != nil {

		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "field validation error"})

	}

	UserUuid := uuid.New()

	NewSportmen.UserId = UserUuid

	JsonSportmen, _ := json.Marshal(NewSportmen)

	req_sport, err := http.NewRequest(http.MethodPost, "http://host.docker.internal:6250/", bytes.NewBuffer(JsonSportmen))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	req_sport.Header.Set("Content-Type", "application/json")

	res_sport, err := http.DefaultClient.Do(req_sport)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	body_sport, err := ioutil.ReadAll(res_sport.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res_sport.StatusCode != 201 {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	defer res_sport.Body.Close()

	fmt.Printf("Code: %d\n", res_sport.StatusCode)
	fmt.Printf("Body: %s\n", body_sport)

	return c.Status(201).JSON(fiber.Map{"status": "success", "sportmen": "Sportmen"})
}
