package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{}

func (c *Config) Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
