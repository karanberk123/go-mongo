package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {
}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	if err := godotenv.Load(filenames...); err != nil {
		panic(err)
	}
	return &ConfigImpl{}
}
