package main

import (
	"log"
	"strings"

	"github.com/football-league/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// LoadENV
// func ConnectENV() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal(".env file not loaded properly")
// 	}
// }

func main() {
	// ConnectENV()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	app.Use(logger.New())

	routes.Setup(app)
	log.Fatal(app.Listen(":8000"))
}
