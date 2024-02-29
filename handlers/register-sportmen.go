package handlers

import (
	"fiber-orchestrator/utils"

	"github.com/gofiber/fiber/v2"

	uuid "github.com/google/uuid"
)

func RegisterSportmen(c *fiber.Ctx) error {
	UserUuid := uuid.New()

	res_sport, _ := utils.CreateSportmen(c, UserUuid)

	if res_sport.StatusCode != 201 {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "sportmen": "Sportmen"})
}
