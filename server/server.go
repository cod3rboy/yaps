// Package server provides http server setup and handler to serve image generation requests.
package server

import (
	"strconv"
	"strings"

	"github.com/cod3rboy/yaps/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupAndListen fires up a http server to handle incoming requests for image generation.
//
// For supported image formats, see [SupportedFormats].
func SetupAndListen() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AllowOrigins(),
		AllowMethods: config.AllowMethods(),
	}))
	router := app.Group(config.PathPrefix())
	router.Get("/:format<regex("+strings.Join(SupportedFormats, "|")+")>", HandlerImage)
	app.Listen(config.Host() + ":" + strconv.Itoa(config.Port()))
}
