package controllers

import (
	"fmt"

	models "github.com/football-league/api/models/database"
	types "github.com/football-league/api/models/types"
	"github.com/gofiber/fiber/v2"
)

var db = models.Database()

func PostTeamData(ctx *fiber.Ctx) error {
	players := types.Players{}
	ctx.BodyParser(&players)
	gameJson, err := models.PostGameData(db, players)
	fmt.Println(gameJson)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	return ctx.Status(201).JSON(players)
	// return ctx.Status(400).JSON(types.StatusCode{
	// 	StatusCode: 400,
	// 	Message:    "Can't find your matching game type",
	// })
}
