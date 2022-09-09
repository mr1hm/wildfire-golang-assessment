package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type ConfigError struct {
	Err error
}

func (ce *ConfigError) Error() string {
	return fmt.Sprintf("[ CONFIG ] Error: %v", ce.Err)
}

type config struct {
	Port string `json:"port"`
}

var (
	cfg  *config
	Port string
)

func ReadConfig() error {
	log.Println("Reading config file...")

	file, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		log.Println(err.Error())
		return &ConfigError{
			Err: errors.New(fmt.Sprintf("Read error: %v", err)),
		}
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Println(err.Error())
		return &ConfigError{
			Err: errors.New(fmt.Sprintf("Unmarshal error: %v", err)),
		}
	}

	Port = cfg.Port

	return nil
}
