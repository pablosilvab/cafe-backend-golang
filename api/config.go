package api

import (
	"log"

	"github.com/tkanos/gonfig"
)

const defaultConfigPath = "./config/config.json"

type Configuration struct {
	Port string
}

func LoadConfig() Configuration {
	config := Configuration{}
	err := gonfig.GetConf(defaultConfigPath, &config)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return config
	}
	return config
}
