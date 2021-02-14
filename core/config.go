package core

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {}

// NewConfig creates and set application config instance
func NewConfig() {
	 appConfig = &AppConfig{}
}

// String returns a string value of key
func (*AppConfig) String(key string) string {
	return viper.GetString(key)
}

// DefaultString returns a string value of key, if not exists key then returns a default value
func (*AppConfig) DefaultString(key string, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}
	return v
}

// Int returns a int value of key
func (*AppConfig) Int(key string) int {
	return viper.GetInt(key)
}

// DefaultInt returns a int value of key, if not exists key then returns a default value
func (*AppConfig) DefaultInt(key string, defaultVal int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

// Bool returns a bool value of key
func (*AppConfig) Bool(key string) bool {
	return viper.GetBool(key)
}

// Duration returns a time.duration value of key
func (*AppConfig) Duration(key string) time.Duration {
	return viper.GetDuration(key)
}

// IsDev returns true if GO_ENV has a development value, otherwise false
func (c *AppConfig) IsDev() bool {
	return c.GetMode() == "development"
}

// GetMode returns GO_ENV value
func (c *AppConfig) GetMode() string {
	mode := c.DefaultString("GO_ENV", "development")
	return mode
}