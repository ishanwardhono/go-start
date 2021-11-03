package config

import (
	"app/env"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

const (
	filePath = "env/env.%s.yaml"
)

var (
	config *Config
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	MaxGraceStop int
	LogFile      string
	LogLevel     string
}

func GetConfig() *Config {
	if config == nil {
		err := NewConfig()
		if err != nil {
			panic("Error get config, err " + err.Error())
		}
	}
	return config
}

func NewConfig() error {
	Cfg := Config{}
	err := readConfigFile(&Cfg, fmt.Sprintf(filePath, env.GetEnv()))
	if err != nil {
		return err
	}
	config = &Cfg
	return nil
}

func readConfigFile(cfg *Config, filename string) error {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		log.Fatal("error in umarshalling env yaml file")
		return err
	}
	return nil
}
