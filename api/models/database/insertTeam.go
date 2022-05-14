package models

import (
	"database/sql"
	"fmt"
	"log"

	types "github.com/football-league/api/models/types"

	_ "github.com/mattn/go-sqlite3"
)

func PostGameData(db *sql.DB, players types.Players) ([]types.Players, error) {
	insert, err := db.Prepare("INSERT INTO players(team_a ,team_b,score_a,score_b,date_of_match) VALUES(?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	_, err = insert.Exec(&players.TeamA, &players.TeamB, &players.ScoreA, &players.ScoreB, &players.DateOfMatch)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println("Sucessfully Data Inserted on users Table")
	return nil, err
}
