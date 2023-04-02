package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"path"
	"sync"
)

type Config struct {
	Kafka struct {
		Server  string `yaml:"server"`
		GroupId string `yaml:"group_id"`
	}
	Mongo struct {
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("parsing yaml config file")
		instance = &Config{}

		err := cleanenv.ReadConfig(path.Base("config.yaml"), instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalf("error due config %v", err)
		}
	})
	return instance
}
