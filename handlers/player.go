// Package handlers
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/bship/dto"
	"github.com/kolakdd/bship/entity"
)

func InitPlayer(c *fiber.Ctx) error {

	dto := new(dto.Name)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't find name in body", "data": err})
	}
	player := entity.NewPlayer(dto.Name)
	return c.JSON(fiber.Map{"status": "success", "message": "created", "data": player})
}
