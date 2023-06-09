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
	Password string `yaml:"password"`
	Email    string `yaml:"from_email"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("parsing yaml config file")
		instance = &Config{}

		err := cleanenv.ReadConfig(path.Join("/home", "chechyotka", "projects", "golang_projects", "car_booking_service", "email_sender_microservice", "config.yaml"), instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalf("error due config %v", err)
		}
	})
	return instance
}
