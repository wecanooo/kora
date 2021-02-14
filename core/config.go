package core

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {}

// NewConfig application config instance 생성
func NewConfig() {
	 appConfig = &AppConfig{}
}

// String string 설정 반환
func (*AppConfig) String(key string) string {
	return viper.GetString(key)
}

// DefaultString default string 설정 반환
func (*AppConfig) DefaultString(key string, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}
	return v
}

// Int int 설정 반환
func (*AppConfig) Int(key string) int {
	return viper.GetInt(key)
}

// DefaultInt default int 설정 반환
func (*AppConfig) DefaultInt(key string, defaultVal int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

// Bool bool 설정 반환
func (*AppConfig) Bool(key string) bool {
	return viper.GetBool(key)
}

// Duration time.duration 설정 반환
func (*AppConfig) Duration(key string) time.Duration {
	return viper.GetDuration(key)
}

// IsDev GO_ENV 설정이 development 라면 true, 아니면 false
func (c *AppConfig) IsDev() bool {
	return c.String("GO_ENV") == "development"
}

// GetMode GO_ENV 설정 반환
func (c *AppConfig) GetMode() string {
	mode := c.DefaultString("GO_ENV", "development")
	return mode
}