package util

import "github.com/spf13/viper"

// DB_DRIVER=postgres
// DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
// SERVER_ADDRESS=0.0.0.0:8080

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	return
}
