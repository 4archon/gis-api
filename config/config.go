package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	GisApi					string	`json:"GisApiKey"`
	PostgresUser			string	`json:"postgresUser"`
	PostgresPassword		string	`json:"postgresPassword"`
	PostgresDbName			string	`json:"postgresDbName"`
	PostgresSSL				string	`json:"postgresSSL"`
	PostgresIsolationLevel	string	`json:"postgresIsolationLevel"`
	JwtSecretKey			[]byte	`json:"jwtSecretKey"`
	AllDataSecretKey		string	`json:"AllDataSecretKey"`
}

func (c *Config) Init() {
	file, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(file, c)
	if err != nil {
		log.Println(err.Error())
		return
	}
}