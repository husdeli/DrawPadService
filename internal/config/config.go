package config

import (
	"sync"

	"github.com/husdeli/DrawPadService.git/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port        int    `env:"PORT" env-required:"true"`
	Host        string `env:"HOST" env-required:"true"`
	IsDev       bool   `env:"IS_DEV" env-required:"true"`
	CacheDbHost string `env:"CACHE_DB_HOST" env-required:"true"`
	CacheDbPort int    `env:"CACHE_DB_PORT" env-required:"true"`
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
	})
	return instance
}
