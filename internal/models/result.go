package models

type GameState int

const (
	Undefined GameState = iota // 0
	Win                        // 1
	Loss                       // 2
	Draw                       // 3
)

type Result struct {
	PlayerOneId    string `json:"player_one_id"`
	PlayerOneState string `json:"player_one_state"`
	PlayerTwoId    string `json:"player_two_id"`
	PlayerTwoState string `json:"player_two_state"`
}
