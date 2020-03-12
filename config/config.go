package config

import (
	"github.com/spf13/viper"
)

func InitConfig(dir string) {
	workDir := dir
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("load config err" + err.Error())
	}
}
