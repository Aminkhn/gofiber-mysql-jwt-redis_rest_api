package config

import "github.com/spf13/viper"

type Configuration struct {
	// DataBase Setup
	DBHost     string `mapstructure:"DB_HOST"`
	DBUserame  string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	// Redis Setup
	RedisUrl string `mapstructure:"REDIS_URL"`
	// Server Setting
	ServerUrl   string `mapstructure:"SERVER_URL"`
	SERVER_Port string `mapstructure:"SERVER_PORT"`
	// jwt secret
	Secret string `mapstructure:"SECRET"`
}

func LoadConfig(path string) (config Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// handle null
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
