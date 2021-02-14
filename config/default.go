package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultTempDir = "tmp"
	defaultAppPort = "3000"
	defaultAppName = "app"
	defaultStorage = "storage"
)

var defaultConfigMap = map[string]interface{}{
	// app
	"APP.NAME":          defaultAppName,
	"APP.VERSION":       "1.0.0",
	"APP.RUNMODE":       "production",
	"APP.ADDR":          ":" + defaultAppPort,
	"APP.URL":           "http://localhost:" + defaultAppPort,
	"APP.KEY":           "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
	"APP.TEMP_DIR":      defaultTempDir,
	"APP.UPLOAD_DIR":    defaultStorage,

	// db
	"DB.DEFAULT.CONNECTION":           "postgres",
	"DB.DEFAULT.HOST":                 "127.0.0.1",
	"DB.DEFAULT.PORT":                 "5432",
	"DB.DEFAULT.DATABASE":             defaultAppName,
	"DB.DEFAULT.USERNAME":             "root",
	"DB.DEFAULT.PASSWORD":             "secret",
	"DB.DEFAULT.OPTIONS":              "",
	"DB.DEFAULT.MAX_OPEN_CONNECTIONS": 100,
	"DB.DEFAULT.MAX_IDLE_CONNECTIONS": 20,

	// jwt token
	"TOKEN.ACCESS_TOKEN_LIFETIME": 60 * time.Minute,
	"TOKEN.REFRESH_TOKEN_LIFETIME": 60 * 24 * time.Minute,

	// log
	"LOG.PREFIX":     "[ZAP_LOGGER]",
	"LOG.FOLDER":     defaultTempDir + "/logs/zap",
	"LOG.LEVEL":      "debug",
	"LOG.MAXSIZE":    10,
	"LOG.MAXBACKUPS": 5,
	"LOG.MAXAGES":    30,
}

func setupDefaultConfig() {
	for k, v := range defaultConfigMap {
		viper.SetDefault(k, v)
	}
}
