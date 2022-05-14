package models

import (
	"database/sql"
	"fmt"

	"github.com/football-league/api/models/types"
	_ "github.com/mattn/go-sqlite3"
)

func GetGames(db *sql.DB, id int, gamePlayerRes types.Players) ([]types.Players, error) {
	var games []types.Players
	query := fmt.Sprintf("SELECT DISTINCT id,team_a,team_b,score_a,score_b,date_of_match FROM players WHERE id = %d;", id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return games, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&gamePlayerRes.Id, &gamePlayerRes.TeamA, &gamePlayerRes.TeamB, &gamePlayerRes.ScoreA, &gamePlayerRes.ScoreB, &gamePlayerRes.DateOfMatch)
		if err != nil {
			fmt.Println(err)
			return games, err
		}
		gamePlayerInfo := types.Players{
			Id:          gamePlayerRes.Id,
			TeamA:       gamePlayerRes.TeamA,
			TeamB:       gamePlayerRes.TeamB,
			ScoreA:      gamePlayerRes.ScoreA,
			ScoreB:      gamePlayerRes.ScoreB,
			DateOfMatch: gamePlayerRes.DateOfMatch,
		}
		games = append(games, gamePlayerInfo)
	}
	return games, err
}
