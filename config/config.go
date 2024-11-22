package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	env *Enviroment
)

func SetConfig() *Config {
	runtime_viper := viper.New()
	runtime_viper.SetConfigName(".env")
	runtime_viper.SetConfigType("env")
	runtime_viper.AddConfigPath(".")
	if err := runtime_viper.ReadInConfig(); err != nil {
		logrus.New().Errorf("Failed to ReadRemoteConfig : %s", err)
	}
	//unmarshal config
	if err := runtime_viper.UnmarshalExact(&env); err != nil {
		logrus.New().Errorf("Failed to UnmarshalExact : %s", err)
	}
	return &Config{Env: env}
}
