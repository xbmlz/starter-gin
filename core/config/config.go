package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xbmlz/starter-gin/utils"
)

type Log struct {
	Path string
}

type Server struct {
	RunMode string
	Address string
	Port    int
}

type AppConfig struct {
	Server Server
	Log    Log
}

var App = &AppConfig{}

// Setup global config, see https://github.com/spf13/viper
func Setup() {
	config := utils.GetEnvString("VIPER_PATH", "config.yaml")
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file is changed:", e.Name)
		if err = v.Unmarshal(&App); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&App); err != nil {
		fmt.Println(err)
	}
}
