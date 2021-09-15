package initialize

import (
	"github.com/akazwz/file-server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitViper() (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
		return nil
	}

	if err := config.Unmarshal(&global.CFG); err != nil {
		panic(err)
		return nil
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file updated:", e.Name)
		if err := config.Unmarshal(&global.CFG); err != nil {
			panic(err)
		}
	})

	return
}
