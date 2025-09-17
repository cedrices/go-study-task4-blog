package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
	jwt    Jwt          `yaml:"jwt"`
}

type DBConfig struct {
	name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

var (
	config Config
	once   sync.Once
)

func LoadConfig() (c *Config, e error) {
	once.Do(func() {
		// Configuration loading logic here
		data, err := os.ReadFile("config/conf.yaml")
		if err != nil {
			fmt.Println("Error reading YAML file:", err)
		}

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			fmt.Println("Error parsing YAML file:", err)
		}
	})
	return &config, nil
}
