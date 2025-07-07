package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/bship/storage"
)

type DebugHandler struct {
	storage *storage.Storage
}

func NewDebugHandler(storage *storage.Storage) *DebugHandler {
	return &DebugHandler{storage: storage}
}

// GetStorageInfo Init game room
func (h *DebugHandler) GetStorageInfo(c *fiber.Ctx) error {
	fmt.Println("init room storage", h.storage)

	return c.JSON(fiber.Map{
		"storage.TokenPlayer":   h.storage.TokenPlayer,
		"storage.Clients":       h.storage.ClientsWs,
		"storage.ClientsPairs":  h.storage.ClientsPairs,
		"Sessions":              h.storage.Sessions,
		"storage.InviteSession": h.storage.InviteSession,
	})
}
