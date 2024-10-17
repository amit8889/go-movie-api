package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to configuration files")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatal(err)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can't read config file: %s", err.Error())
	}
	return &cfg

}
