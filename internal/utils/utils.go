package utils

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetupConfig() {
	viper.SetConfigName("init-app")
	viper.SetConfigType("yaml")
	homeDir, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDir, "init-app.yaml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if _, err := os.Create(configFile); err != nil {
			panic(err)
		}
	}
	viper.AddConfigPath(homeDir)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
