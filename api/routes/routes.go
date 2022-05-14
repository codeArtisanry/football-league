package routes

import (
	controllers "github.com/football-league/api/controllers"
	"github.com/gofiber/fiber/v2"
)

// Routes for diffrent Endpoints
func Setup(app *fiber.App) {
	app.Post("api/v1/players", controllers.PostTeamData)
	app.Get("api/v1/players", controllers.GetGames)
	app.Get("api/v1/players/:id", controllers.GetGames)
}
