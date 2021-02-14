package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wecanooo/kora/core"
	"github.com/wecanooo/kora/core/pkg/db"
)

func SetupDefaultDatabase() *sql.DB {
	dsn := db.BuildDatabaseDSN(core.GetConfig().DefaultString("DB.DEFAULT.CONNECTION", "postgres"), db.DatabaseConfig{
		UserName: core.GetConfig().String("DB.DEFAULT.USERNAME"),
		Password: core.GetConfig().String("DB.DEFAULT.PASSWORD"),
		Host:     core.GetConfig().String("DB.DEFAULT.HOST"),
		Port:     core.GetConfig().String("DB.DEFAULT.PORT"),
		DBName:   core.GetConfig().String("DB.DEFAULT.DATABASE"),
		Options:  core.GetConfig().String("DB.DEFAULT.OPTIONS"),
	}, func(s string) string {
		return core.GetConfig().String("DB.DEFAULT.DATABASE") + "_" + core.GetConfig().GetMode()
	})

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Database 설정 중 오류가 발생되었습니다, 오류 :" + err.Error())
	}
	conn.SetMaxOpenConns(core.GetConfig().DefaultInt("DB.DEFAULT.MAX_OPEN_CONNECTIONS", 100))
	conn.SetMaxIdleConns(core.GetConfig().DefaultInt("DB.DEFAULT.MAX_IDLE_CONNECTIONS", 20))

	fmt.Printf("\nDefault atabase connection successful: %s\n", dsn)
	return conn
}
