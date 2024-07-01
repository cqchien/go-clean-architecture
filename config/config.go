package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port       string `env:"PORT" envDefault:"8080"`
	HashSalt   string `env:"HASH_SALT,required"`
	SigningKey string `env:"SIGNING_KEY,required"`
	TokenTTL   int64  `env:"TOKEN_TTL,required"`
	JwtSecret  string `env:"JWT_SECRET,required"`
	DbHost     string `env:"DB_HOST,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbPort     int64  `env:"DB_PORT,required"`
	DbName     string `env:"DB_NAME,required"`
}

func GetConfig(fileNames ...string) *Configuration {
	err := godotenv.Load(fileNames...)

	if err != nil {
		log.Println("No .env file could be found ", fileNames)
	}

	config := Configuration{}

	errorParseEnv := env.Parse(&config)

	if errorParseEnv != nil {
		log.Println("Can not parse env data. Error: ", errorParseEnv)
	}

	return &config
}
