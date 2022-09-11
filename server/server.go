package server

import (
	"strconv"
	"strings"

	"github.com/cod3rboy/yaps/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.GetAllowOrigins(),
		AllowMethods: config.GetAllowMethods(),
	}))
	router.Get("/:format<regex("+strings.Join(SupportedFormats, "|")+")>", HandlerImage)
	router.Listen(config.GetHost() + ":" + strconv.Itoa(config.GetPort()))
}
