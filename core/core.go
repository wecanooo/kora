package core

import (
	"github.com/go-redis/redis"
	db "github.com/wecanooo/kora/app/models"
	"go.uber.org/zap"
)

var (
	application       *Application
	defaultConnection *DBConn
	store			  db.Store
	appConfig         *AppConfig
	redisClient 	  *redis.Client
	appLog            *zap.SugaredLogger
)

// GetApplication return a application instance
func GetApplication() *Application {
	if application == nil {
		panic("application is not initialized")
	}
	return application
}

// GetStore returns a application store interface
func GetStore() db.Store {
	if store == nil {
		panic("store is not initialized")
	}
	return store
}

// GetDefaultConnection returns a database connection
func GetDefaultConnection() *DBConn {
	if defaultConnection == nil {
		panic("database connection is not initialized")
	}
	return defaultConnection
}

// GetConfig application config 반환
func GetConfig() *AppConfig {
	if appConfig == nil {
		panic("application config is not initialized")
	}
	return appConfig
}

// GetLog returns a application logger instance
func GetLog() *zap.SugaredLogger {
	if appLog == nil {
		panic("application logger is not initialized")
	}
	return appLog
}

// GetRedisClient returns a redis client instance
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		panic("redis client is not initialized")
	}
	return redisClient
}
