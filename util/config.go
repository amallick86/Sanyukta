package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CountryStateAPIToken string        `mapstructure:"COUNTRY_STATE_API_TOKEN"`
	SendInBlueKey        string        `mapstructure:"SEND_IN_BLUE_KEY"`
	SendInBlueEmail      string        `mapstructure:"SEND_IN_BLUE_EMAIL"`
	SenderEmail          string        `mapstructure:"SENDER_EMAIL"`
	SendInBluePassword   string        `mapstructure:"SEND_IN_BLUE_PASSWORD"`
	SMTPHost             string        `mapstructure:"SMTP_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
