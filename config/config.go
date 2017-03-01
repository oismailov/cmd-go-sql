package config

import (
	"encoding/json"
	"log"
	"os"
)

var Cfg *Config = &Config{}

//Config struct
type Config struct {
	DatabaseSettings DatabaseSettings
}

//DatabaseSettings of the project
type DatabaseSettings struct {
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
}

//LoadConfig from conf/conf.json  file
func LoadConfig() {
	file, errOpenFile := os.Open("config/conf.json")
	if errOpenFile != nil {
		log.Fatal(errOpenFile)
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("error:", err)
	}
	Cfg = &configuration
}
