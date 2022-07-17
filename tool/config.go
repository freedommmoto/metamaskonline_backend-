package tool

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type ConfigObject struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOUECE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenConfigKey      string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	PusherKey           string        `mapstructure:"PUSHER_KEY"`
	TokenLiftTimeConfig time.Duration `mapstructure:"ACCESS_TOKEN_DURATION_TIME"`
}

func LoadConfig(part string) (config ConfigObject, err error) {
	viper.AddConfigPath(part)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct %v", err)
	}
	return
}
