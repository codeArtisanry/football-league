package controllers

import (
	"fmt"

	types "github.com/foot-league/api/models/types"
	models "github.com/football-league/api/models/database"
	"github.com/gofiber/fiber/v2"
)

var db = models.Database()

func PostTeamData(ctx *fiber.Ctx) error {
	fmt.Println("hello")
	UserInfo := types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
	id, err := models.VerifyAndInsertUser(db, UserInfo)
	return nil
}
