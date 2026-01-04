package models

import (
	"time"
)

type User struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	Total_win   int32  `db:"total_win"`
	Total_games int32  `db:"total_games"`

	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}
