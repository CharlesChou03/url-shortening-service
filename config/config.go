package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type envConfig struct {
	Version              string `env:"VERSION" envDefault:"0.0.1"`
	MongoURI             string `env:"MONGO_URI,required"`
	ShorteningUrlLength  int    `env:"SHORTENING_URL_LENGTH,required"`
	UrlHost              string `env:"URL_HOST,required"`
	DefaultExpiredPeriod int64  `env:"DEFAULT_EXPIRED_PERIOD,required"`
}

var (
	// Env is the config
	Env = envConfig{}
)

// Setup setup config function
func Setup() {
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%+v\n", Env)
}
