package config

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/google/wire"

	"gopkg.in/yaml.v2"
)

type IDatabaseConfig interface {
	GetHost() string
	GetPort() uint
	GetUsername() string
	GetPassword() string
	GetDatabase() string
	GetMaxOpenConnections() int
	GetMaxIdleConnections() int
}

type Configuration struct {
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

type AppConfiguration struct {
	*Configuration
}

var NewAppConfiguration = wire.NewSet(
	GetConfiguration,
	wire.Struct(new(AppConfiguration), "*"),
	wire.Bind(new(IDatabaseConfig), new(*AppConfiguration)),
)

var (
	once          sync.Once
	configuration *Configuration
)

func GetConfiguration() *Configuration {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Fatalf("failed loading configuration: %v", err)
		}
		err = yaml.Unmarshal(yamlFile, configuration)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	})
	return configuration
}

func (ac *AppConfiguration) GetHost() string {
	return ac.DB.Host
}

func (ac *AppConfiguration) GetPort() uint {
	return ac.DB.Port
}

func (ac *AppConfiguration) GetUsername() string {
	return ac.DB.Username
}

func (ac *AppConfiguration) GetPassword() string {
	return ac.DB.Password
}

func (ac *AppConfiguration) GetDatabase() string {
	return ac.DB.Database
}

func (ac *AppConfiguration) GetMaxOpenConnections() int {
	return ac.DB.MaxOpenConnections
}

func (ac *AppConfiguration) GetMaxIdleConnections() int {
	return ac.DB.MaxIdleConnections
}
