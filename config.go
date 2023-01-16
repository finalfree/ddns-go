package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	CorpID    string `yaml:"corp_id"`
	AppSecret string `yaml:"app_secret"`
	AgentID   int    `yaml:"agent_id"`
	ToUser    string `yaml:"to_user"`
}

type TempConfig struct {
	Token    string `yaml:"token"`
	ExpireAt int64  `yaml:"expire_at"`
	LastIP   string `yaml:"last_ip"`
}

func ReadConfig(fileName string) *Config {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// read yaml config from file
	decoder := yaml.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return &config
}

func ReadTempConfig(fileName string) (*TempConfig, *os.File) {
	var file *os.File
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(fileName), 0770); err != nil {
			fmt.Println(err)
		}

		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		file, err = os.OpenFile(fileName, os.O_RDWR, 0660)
		if err != nil {
			fmt.Println(err)
		}
	}

	decoder := yaml.NewDecoder(file)
	var config TempConfig
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return &config, file
}

func WriteTempConfig(file *os.File, config *TempConfig) error {
	encoder := yaml.NewEncoder(file)
	err := encoder.Encode(config)
	return err
}
