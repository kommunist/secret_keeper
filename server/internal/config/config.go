package config

import (
	"flag"
	"log/slog"
	"os"
)

type MainConfig struct {
	DatabaseURI string
	ServerURL   string
	CertPath    string
	CertKeyPath string
}

func Make() MainConfig {
	config := MainConfig{
		DatabaseURI: "",
		ServerURL:   "",
		CertPath:    "certs/MyCertificate.crt",
		CertKeyPath: "certs/MyKey.key",
	}

	return config
}

func (c *MainConfig) Init() {
	c.ParseEnv()
	c.InitFlags()
}

func (c *MainConfig) InitFlags() {
	flag.StringVar(&c.ServerURL, "a", c.ServerURL, "server uri")
	flag.StringVar(&c.DatabaseURI, "d", c.DatabaseURI, "database uri")

	flag.Parse()
	slog.Info("database uri", "uri", c.DatabaseURI)
}

func (c *MainConfig) ParseEnv() {
	if e := os.Getenv("SERVER_URI"); e != "" {
		c.ServerURL = e
	}
	if e := os.Getenv("DATABASE_URI"); e != "" {
		c.DatabaseURI = e
	}

	slog.Info("env parsed")
}
