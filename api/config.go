package api

import (
	"log"

	"github.com/tkanos/gonfig"
)

const defaultConfigPath = "./config/config.json"

// Configuration: Struct for get config information from config.json file.
type Configuration struct {
	Port string
}

// LoadConfig: Function for load config information
func LoadConfig() Configuration {
	config := Configuration{}
	err := gonfig.GetConf(defaultConfigPath, &config)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return config
	}
	return config
}
