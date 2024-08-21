package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	UserSvcUrl         string `mapstructure:"USER_SVC_URL"`
	NotificationSvcUrl string `mapstructure:"NOTIFICATION_SVC_URL"`
	TaskSvcUrl         string `mapstructure:"TASK_SVC_URL"`
	DBUrl              string `mapstructure:"DB_URL"`
	KVAddr             string `mapstructure:"REDIS_ADDR"`
}

func NewConfig() (Config, error) {
	var cfg Config

	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
