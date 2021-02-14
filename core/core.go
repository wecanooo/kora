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

// GetApplication application instance 반환
func GetApplication() *Application {
	if application == nil {
		panic("application is not initialized")
	}
	return application
}

// GetConfig application config 반환
func GetConfig() *AppConfig {
	if appConfig == nil {
		panic("application config is not initialized")
	}
	return appConfig
}

// GetLog application logger 반환
func GetLog() *zap.SugaredLogger {
	if appLog == nil {
		panic("application logger is not initialized")
	}
	return appLog
}
