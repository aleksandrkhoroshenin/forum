package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

const (
	localSettingsFile = "forum-settings.json"
)

//easyjson:json
type Config struct {
	Port       int    `json:"forum.service.port"`
	DbHost     string `json:"forum.service.db.host"`
	DbPort     int    `json:"forum.service.db.port"`
	DbDatabase string `json:"forum.service.db.database"`
	DbUser     string `json:"forum.service.db.user"`
	DbPassword string `json:"forum.service.db.password"`
}

var ToolConfig *Config

func GetInstance() *Config {
	return ToolConfig
}

func readConfig(fileName string) error {
	configFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New("Can't open properties file: " + err.Error())
	}
	if err = ToolConfig.UnmarshalJSON(configFile); err != nil {
		return errors.New("Can't parsing properties file: " + err.Error())
	}
	return nil
}

func InitConfig() error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	configFileName := path.Join(dir, localSettingsFile)
	ToolConfig = new(Config)
	if err := readConfig(configFileName); err != nil {
		return err
	}

	return nil
}
