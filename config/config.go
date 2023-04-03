package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Database struct {
		DbUri       string `yaml:"uri"`
		DbName      string `yaml:"name"`
		CollBoards  string `yaml:"collection_boards"`
		CollThreads string `yaml:"collection_threads"`
		CollPosts   string `yaml:"collection_posts"`
	} `yaml:"database"`
	Server struct {
		ServerHost string `yaml:"host"`
		ServerPort string `yaml:"port"`
	} `yaml:"server"`
}

func NewConfig() *Config {

	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
