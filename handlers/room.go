// Package handlers
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/bship/dto"
	"github.com/kolakdd/bship/entity"
	"github.com/kolakdd/bship/storage"
)

type RoomHandler struct {
	storage *storage.Storage
}

func NewRoomHandler(storage *storage.Storage) *RoomHandler {
	return &RoomHandler{storage: storage}
}

// InitRoom initializate game room
func (h *RoomHandler) InitRoom(c *fiber.Ctx) error {
	dto := new(dto.Name)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Bad req", "data": err})
	}
	player := entity.NewPlayer(dto.Name)
	player.ItLeftPlayer = true

	session := h.storage.CreateSession(&player)
	h.storage.AddTokenPlayer(&player)

	return c.JSON(fiber.Map{"status": "success", "message": "created", "data": fiber.Map{"player": player, "session": session}})
}

// JoinRoom join player in room by invite code
func (h *RoomHandler) JoinRoom(c *fiber.Ctx) error {
	dto := new(dto.NameRoom)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Bad req", "data": err})
	}
	player := entity.NewPlayer(dto.Name)

	if err := h.storage.JoinToSession(&player, dto.InviteCode); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "room not found", "data": err})
	}
	h.storage.AddTokenPlayer(&player)

	return c.JSON(fiber.Map{"status": "success", "message": "created", "data": fiber.Map{"player": player}})
}
