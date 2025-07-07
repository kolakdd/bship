// Package router
package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kolakdd/bship/appctx"
	"github.com/kolakdd/bship/handlers"
	wsHandler "github.com/kolakdd/bship/handlers/websocket"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App, appCtx *appctx.Application) {

	// Api
	api := app.Group("/api", logger.New())

	debugHandler := handlers.NewDebugHandler(appCtx.Storage)
	api.Get("/debug", debugHandler.GetStorageInfo)
	// Room
	roomHandler := handlers.NewRoomHandler(appCtx.Storage)
	room := api.Group("/room")
	room.Post("/init", roomHandler.InitRoom)
	room.Post("/join", roomHandler.JoinRoom)
	// Ws
	wsHandler := wsHandler.NewWsHandler(appCtx.Storage)
	ws := app.Group("/ws", logger.New())
	ws.Get("/:token", websocket.New(wsHandler.GameHandler))
}
