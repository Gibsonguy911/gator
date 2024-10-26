package config

import (
	"encoding/json"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var config Config

	home, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	file, err := os.Open(home + "/" + configFileName)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (config Config) SetUser(name string) error {
	config.CurrentUserName = name
	return write(config)
}

func write(config Config) error {
	// Open the file
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	file, err := os.Create(home + "/.gatorconfig.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Marshal the Config struct into JSON
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
