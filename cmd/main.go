package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kolakdd/bship/appctx"
	"github.com/kolakdd/bship/router"
	"github.com/kolakdd/bship/storage"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	// init appContext
	storage := storage.New()
	appCtx := &appctx.Application{
		Storage: storage,
	}
	// todo: ws closer proc
	// go ...

	// main app
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Bship",
	})
	router.SetupRoutes(app, appCtx)
	log.Fatal(app.Listen(":3000"))
}
