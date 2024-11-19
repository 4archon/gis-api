package config

import (
	"os"
)

type Config struct {
	GisApi string
}

func (c *Config) initGis() {
	file, err := os.ReadFile("config/config")
	if err != nil {
		return
	}
	s := string(file)
	c.GisApi = s
}

func (c *Config) Init() {
	c.initGis()
}