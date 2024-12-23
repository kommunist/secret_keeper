package config

import (
	"flag"
	"os"
	"secret_keeper/internal/client/logger"
)

type MainConfig struct {
	DatabaseURI string
	ServerURI   string
}

func Make() MainConfig {
	config := MainConfig{
		DatabaseURI: "",
		ServerURI:   "",
	}

	return config
}

func (c *MainConfig) Init() {
	c.ParseEnv()
	c.InitFlags()
}

func (c *MainConfig) InitFlags() {
	flag.StringVar(&c.ServerURI, "a", c.ServerURI, "server uri")
	flag.StringVar(&c.DatabaseURI, "d", c.DatabaseURI, "database uri")

	flag.Parse()
	logger.Logger.Info("database uri", "uri", c.DatabaseURI)
}

func (c *MainConfig) ParseEnv() {
	if e := os.Getenv("SERVER_URI"); e != "" {
		c.ServerURI = e
	}
	if e := os.Getenv("DATABASE_URI"); e != "" {
		c.DatabaseURI = e
	}

	logger.Logger.Info("env parsed")
}
