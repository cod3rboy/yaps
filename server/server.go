package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
	}))
	router.Get("/:format<regex("+strings.Join(SupportedFormats, "|")+")>", HandlerImage)
	router.Listen("localhost:3001")
}
