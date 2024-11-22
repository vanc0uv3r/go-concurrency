package main

import (
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct{
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
	MaxConnections uint16 `yaml:"max_connections"`
	MaxMsgSize int `yaml:"max_msg_size"`
}

func NewConfig() Config{
	return Config{
		Addr: "127.0.0.1",
		Port: "8888",
		MaxConnections: 10,
		MaxMsgSize: 4096,
	}
}

func importConfig() Config{
	var fileConfig Config

	defaultConfig := NewConfig()
	serverConfig, err := os.ReadFile("config/server.yaml")
	if err != nil {
		log.Println(defaultParamsMsg)
		return defaultConfig
	}
	
	err = yaml.Unmarshal(serverConfig, &fileConfig)
    if err != nil {
		log.Println(parseConfigErr)
		return defaultConfig
    }

	return fileConfig
}