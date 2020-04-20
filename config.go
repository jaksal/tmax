package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
type Config struct {
	Transmission struct {
		URL    string `json:"url"`
		UserID string `json:"userid"`
		Passwd string `json:"passwd"`
	} `json:"transmission"`
	Site   string `json:"site"`
	Search []struct {
		Category string     `json:"category"`
		Finds    [][]string `json:"finds"`
		Save     string     `json:"save"`
	} `json:"search"`
}

func loadConfig(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(f, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
