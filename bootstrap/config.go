package bootstrap

import (
	"github.com/wecanooo/kora/config"
	"github.com/wecanooo/kora/core"
)

// SetupConfig prepares a application config
func SetupConfig(configFilePath, configFileType string) {
	config.Setup(configFilePath, configFileType)
	core.NewConfig()
	config.WriteConfig(core.GetConfig().String("APP.TEMP_DIR") + "/config.json")
}