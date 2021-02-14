package bootstrap

import (
	"github.com/wecanooo/kora/core"
	"github.com/wecanooo/kora/database"
)

// SetupDB database 연결설정
func SetupDB() {
	sqlDB := database.SetupDefaultDatabase()
	core.NewDBConn(sqlDB)
}

// SetupRedis redis 연결설정
func SetupRedis() {
	client := database.SetupRedis()
	core.NewRedis(client)
}
