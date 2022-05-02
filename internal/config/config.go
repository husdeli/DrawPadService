package config

import (
	"sync"

	"github.com/husdeli/DrawPadService.git/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port  string `env:"PORT" env-required:"true"`
	Host  string `env:"HOST" env-required:"true"`
	IsDev bool   `env:"IS_DEV" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		var logger = logger.GetLogger()
		instance = &Config{}
		err := cleanenv.ReadConfig(".env", instance)

		if err != nil {
			logger.Info(cleanenv.GetDescription(instance, nil))
			logger.Fatal(err)
		}
		logger.Infof("Read app config %v", instance)
	})
	return instance
}
