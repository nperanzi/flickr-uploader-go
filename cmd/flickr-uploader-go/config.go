package main

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type config struct {
	TokenFileName     string   `yaml:"token_file_name"`
	APIKey            string   `yaml:"api_key"`
	APISecret         string   `yaml:"api_secret"`
	DbPath            string   `yaml:"db_path"`
	PhotosPath        string   `yaml:"photos_path"`
	ExcludeDirs       []string `yaml:"exclude_dirs"`
	APIRequestSleepMs int      `yaml:"api_request_sleep_ms"`
}

// todo возвращать не указатель
func newConfig(fileName string) (*config, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil, errors.WithStack(err)
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	newConfig := config{}
	if err := yaml.Unmarshal(file, &newConfig); err != nil {
		return nil, errors.WithStack(err)
	}
	return &newConfig, nil
}
