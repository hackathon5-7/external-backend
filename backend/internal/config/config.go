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

// MustLoad loads the configuration from the specified file path.
// It returns a pointer to the loaded configuration.
// If the file path is empty or the file does not exist, it panics.
func MustLoad() *Config {
	configFilePath := os.Getenv("CONFIG_PATH")
	if configFilePath == "" {
		log.Panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configFilePath); err != nil {
		log.Panicf("config file %s does not exist", configFilePath)
	}

	config := &Config{}
	if err := cleanenv.ReadConfig(configFilePath, config); err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	return config
}
