package config

import (
	"client/internal/logger"
	"flag"
	"os"
)

type MainConfig struct {
	DatabaseURI string
	ServerURL   string
}

func Make() *MainConfig {
	config := MainConfig{
		DatabaseURI: "",
		ServerURL:   "localhost:1025",
	}

	config.ParseEnv()
	config.InitFlags()

	return &config
}

func (c *MainConfig) InitFlags() {
	flag.StringVar(&c.ServerURL, "a", c.ServerURL, "server uri")
	flag.StringVar(&c.DatabaseURI, "d", c.DatabaseURI, "database uri")

	flag.Parse()
	logger.Logger.Info("database uri", "uri", c.DatabaseURI)
}

func (c *MainConfig) ParseEnv() {
	if e := os.Getenv("SERVER_URI"); e != "" {
		c.ServerURL = e
	}
	if e := os.Getenv("DATABASE_URI"); e != "" {
		c.DatabaseURI = e
	}

	logger.Logger.Info("env parsed")
}
