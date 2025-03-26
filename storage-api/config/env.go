package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBName               string `mapstructure:"DB_NAME"`
	WebServerPort        string `mapstructure:"PORT"`
	AmqpServerUrl        string `mapstructure:"AMQP_SERVER_URL"`
	QueueNameFromBilling string `mapstructure:"BILLING_TO_STORAGE_QUEUE"`
	QueueNameToBilling   string `mapstructure:"STORAGE_TO_BILLING_QUEUE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
