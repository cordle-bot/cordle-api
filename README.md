# **`cordle-api`**

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/cordle-bot/cordle-api?style=flat-square)

API for Cordle Bot

## *`Docs`*

<https://pkg.go.dev/github.com/cordle-bot/cordle-api>

## *`Modules`*

- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm)

## *`Usage`*

### *`Running`*

1. `$ go mod download`

2. Ensure you have a `.env` in the project root following `.env.example` as a guide.

3. `$ make run`

### *`Docker`*

TODO: dockerise

## *`Models`*

```go
type UserModel struct {
    Id     string `gorm:"primaryKey;not null"`           // primary key
    Wins   int    `gorm:"default:0;type:int;not null"`   // wins, default 0, int, not null
    Losses int    `gorm:"default:0;type:int;not null"`   // losses, default 0, int, not null
    Draws  int    `gorm:"default:0;type:int;not null"`   // draws, default 0, int, not null
    Elo    int    `gorm:"default:500;type:int;not null"` // elo, default 500, int, not null
}
```
