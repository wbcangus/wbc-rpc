package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
)

type Config struct {
	Rpc struct {
		Name    string `yaml:"name"`
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Version string `yaml:"version"`
	} `yaml:"rpc"`
	Sys struct {
		LogLevel string `yaml:"logLevel"`
	} `yaml:"system"`
}

var cfg *Config
var once sync.Once

func loadConfig() error {
	file, err := os.Open("./resource/config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}(file)
	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	cfg = &config
	return nil
}

func GetConfig() *Config {
	once.Do(func() {
		err := loadConfig()
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
