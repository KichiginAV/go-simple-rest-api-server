package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ListenHTTP string `yaml:"listen_http"`
	DataBase   DB     `yaml:"database"`
}

type DB struct {
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	DBName   string `yaml:"db"`
	Driver   string `yaml:"driver"`
}

func Init(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %v", err)
	}

	conf := new(Config)
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, fmt.Errorf("не удалось разобрать файл конфигурации: %v", err)
	}

	if conf.ListenHTTP == "" {
		conf.ListenHTTP = ":7777"
	}
	return conf, nil
}
