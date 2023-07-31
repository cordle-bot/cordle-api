package models

type Leaderboard struct {
	TopTen [10]User `json:"top_ten"`
}
