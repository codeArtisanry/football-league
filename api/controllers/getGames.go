package controllers

import (
	"strconv"

	models "github.com/football-league/api/models/database"
	types "github.com/football-league/api/models/types"
	"github.com/gofiber/fiber/v2"
)

func GetGames(ctx *fiber.Ctx) error {
  
	gameId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	gameres := types.Players{}
	gamebyid, err := models.GetGames(db, gameId, gameres)
	if err != nil {
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	return ctx.Status(200).JSON(gamebyid)

}
