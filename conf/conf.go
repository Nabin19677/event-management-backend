package conf

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	DatabaseSource string `mapstructure:"DATABASE_SOURCE"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	JwtSecret      string `mapstructure:"JWT_SECRET"`
	JwtIssuer      string `mapstructure:"JWT_ISSUER"`
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
