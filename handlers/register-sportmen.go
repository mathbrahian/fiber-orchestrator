package handlers

import (
	"fiber-orchestrator/utils"

	"github.com/gofiber/fiber/v2"

	uuid "github.com/google/uuid"
)

type response struct {
	data       []byte
	statusCode int
	err        error
}

func GoFuncion(c *fiber.Ctx, UserUuid uuid.UUID) (response, response) {
	response_sport := make(chan response)
	response_user := make(chan response)
	go func() {
		data, statusCode, err := utils.CreateSportmen(c, UserUuid)
		response_sport <- response{data, statusCode, err}
	}()
	go func() {
		data, statusCode, err := utils.CreateUser(c, UserUuid)
		response_user <- response{data, statusCode, err}
	}()
	return <-response_sport, <-response_user
}

func RegisterSportmen(c *fiber.Ctx) error {
	UserUuid := uuid.New()
	response_sport, response_user := GoFuncion(c, UserUuid)

	if response_sport.err != nil && response_user.err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	if response_sport.statusCode != 201 && response_user.statusCode != 201 {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "sportmen": "Sportmen"})
}
