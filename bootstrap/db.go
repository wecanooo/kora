package bootstrap

import (
	db "github.com/wecanooo/kora/app/models"
	"github.com/wecanooo/kora/core"
	"github.com/wecanooo/kora/database"
)

// SetupDB prepares database connection and store
func SetupDB() {
	sqlDB := database.SetupDefaultDatabase()
	store := db.NewSQLStore(sqlDB)
	core.NewDBConn(sqlDB)
	core.NewStore(store)
}

// SetupRedis prepares a redis client
func SetupRedis() {
	client := database.SetupRedis()
	core.NewRedis(client)
}
