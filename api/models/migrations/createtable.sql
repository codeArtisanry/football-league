-- +migrate Up
CREATE TABLE IF NOT EXISTS players(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_a TEXT,
    team_b TEXT,
		score_a INTEGER,
		score_b INTEGER,
		date_of_match DATETIME,
       created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    )
);
-- +migrate Down
DROP TABLE IF EXISTS players;
