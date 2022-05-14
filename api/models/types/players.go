package types

import "time"

type Players struct {
	Id          int        `json:"id"`
	TeamA       string     `json:"team_a"`
	TeamB       string     `json:"team_b"`
	ScoreA      int        `json:"score_a"`
	ScoreB      int        `json:"score_b"`
	DateOfMatch *time.Time `json:"date_of_match"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type StatusCode struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}
