package models

type UserModel struct {
	Id     string `gorm:"primaryKey;not null"`           // primary key
	Wins   int    `gorm:"default:0;type:int;not null"`   // wins, default 0, int, not null
	Losses int    `gorm:"default:0;type:int;not null"`   // losses, default 0, int, not null
	Draws  int    `gorm:"default:0;type:int;not null"`   // draws, default 0, int, not null
	Elo    int    `gorm:"default:500;type:int;not null"` // elo, default 500, int, not null
}

func (u *UserModel) ToUserPost() UserPost {
	return UserPost{
		Id:     u.Id,
		Wins:   u.Wins,
		Losses: u.Losses,
		Draws:  u.Draws,
		Elo:    u.Elo,
	}
}

func (u *UserModel) ToUser() User {
	return User{
		Id:     u.Id,
		Wins:   u.Wins,
		Losses: u.Losses,
		Draws:  u.Draws,
		Games:  u.Wins + u.Losses + u.Draws,
		Elo:    u.Elo,
	}
}

type UserPost struct {
	Id     string `json:"id"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Draws  int    `json:"draws"`
	Elo    int    `json:"elo"`
}

func (u *UserPost) ToUserModel() UserModel {
	return UserModel{
		Id:     u.Id,
		Wins:   u.Wins,
		Losses: u.Losses,
		Draws:  u.Draws,
		Elo:    u.Elo,
	}
}

func (u *UserPost) ToUser() User {
	return User{
		Id:     u.Id,
		Wins:   u.Wins,
		Losses: u.Losses,
		Draws:  u.Draws,
		Games:  u.Wins + u.Losses + u.Draws,
		Elo:    u.Elo,
	}
}

type User struct {
	Id     string // user id
	Wins   int    // user wins
	Losses int    // user losses
	Draws  int    // user draws
	Games  int    // user games played
	Elo    int    // user elo
}
