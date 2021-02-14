package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// Setup ...
func Setup(configFilePath, configFileType string) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("프로그램 설정을 로드할 수 없습니다, 설정파일: %s, 오류: %v", configFilePath, err))
	}

	setupDefaultConfig()

	viper.AutomaticEnv()
	viper.SetEnvPrefix(viper.GetString("APP.NAME"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	watchConfig()

	fmt.Printf("\nConfiguration file loaded successfully: %s\n", configFilePath)
}

// WriteConfig ...
func WriteConfig(filename string) {
	viper.WriteConfigAs(filename)
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		log.Infof("Config file changed: %s", ev.Name)
	})
}
