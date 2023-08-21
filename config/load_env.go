package config

import "github.com/spf13/viper"

type Configuration struct {
	// DataBase Setup
	DBHost     string `mapstructure:"SQL_HOST"`
	DBUserame  string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASSWORD"`
	DBName     string `mapstructure:"SQL_DB"`
	DBPort     string `mapstructure:"SQL_PORT"`
	// Redis Setup
	RedisUrl string `mapstructure:"REDIS_URL"`
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
