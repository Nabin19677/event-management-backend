package conf

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.SetConfigType("toml")
	viper.SetConfigName("conf/conf")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Overwrite ENV CONFIGS
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
