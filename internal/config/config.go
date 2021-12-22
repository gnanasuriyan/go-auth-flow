package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		Host               string `default:"127.0.0.1" yaml:"host"`
		Port               uint   `default:"3306" yaml:"port"`
		Username           string `default:"root" yaml:"username"`
		Password           string `default:"password" yaml:"password"`
		Database           string `default:"acme_corp" yaml:"database"`
		MaxOpenConnections int    `default:"50" yaml:"max_open_connections"`
		MaxIdleConnections int    `default:"25" yaml:"max_idle_connections"`
	}
}

type AppConfig struct {
	*Config
}

var (
	configOnce sync.Once
	appConfig  *AppConfig
)

func GetConfig() *AppConfig {
	configOnce.Do(func() {
		yamlFile, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Fatalf("failed loading configuration: %v", err)
		}
		err = yaml.Unmarshal(yamlFile, appConfig)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	})
	return appConfig
}
