package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
)

type Config struct {
	PostgreSQL struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
	App struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Network string `yaml:"network"`
	} `yaml:"app"`
}

const (
	EnvConfigPathName  = "CONFIG_PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(
			&configPath,
			FlagConfigPathName,
			"config.yaml",
			"this is app config file",
		)
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		// Чтение содержимого YAML-файла
		yamlFile, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("Ошибка чтения файла конфигурации: %v", err)
		}

		// Создание экземпляра структуры Config
		instance = &Config{}

		// Анмаршализация YAML в структуру Config
		err = yaml.Unmarshal(yamlFile, &instance)
		if err != nil {
			log.Fatalf("Ошибка разбора YAML: %v", err)
		}
	})
	return instance
}
