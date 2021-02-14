package db

import (
	"fmt"
	"log"
)

// DatabaseConfig ...
type DatabaseConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
	Options  string
}

type connectionName string

// BuildDatabaseDSN 데이터베이스 연결설정
func BuildDatabaseDSN(connection string, config DatabaseConfig, buildDBName func(string) string) string {
	if buildDBName != nil {
		config.DBName = buildDBName(config.DBName)
	}

	switch connection {
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?%s",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
			config.Options,
		)
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s options='%s'",
			config.Host,
			config.Port,
			config.UserName,
			config.DBName,
			config.Password,
			config.Options,
		)
	case "sqlite3":
		return config.DBName
	case "mssql":
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s&%s",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
			config.Options,
		)
	}

	log.Panicf("DB Connection %s not supported", connection)
	return ""
}
