package bootstrap

import (
	db "github.com/wecanooo/kora/app/models"
	"github.com/wecanooo/kora/core"
	"github.com/wecanooo/kora/database"
)

// SetupDB database 연결설정
func SetupDB() {
	sqlDB := database.SetupDefaultDatabase()
	store := db.NewSQLStore(sqlDB)
	core.NewDBConn(sqlDB)
	core.NewStore(store)
}

// SetupRedis redis 연결설정
func SetupRedis() {
	client := database.SetupRedis()
	core.NewRedis(client)
}
