package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer  HTTPServer  `yaml:"http_server"`
	DataBase    DataBase    `yaml:"db"`
	RedisConfig RedisConfig `yaml:"redis"`
}

type HTTPServer struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Timeout     string `yaml:"timeout"`
	IdleTimeout string `yaml:"idle_timeout"`
}

type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"password"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	TTL  int    `yaml:"ttl"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH IS NOT SET")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("not found %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("%s", err)
	}

	return &cfg
}
