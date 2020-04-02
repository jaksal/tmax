package main

// Config ...
type Config struct {
	Transmission struct {
		URL    string `json:"url"`
		UserID string `json:"userid"`
		Passwd string `json:"passwd"`
	} `json:"transmission"`
	Search []struct {
		Category string     `json:"category"`
		Finds    [][]string `json:"finds"`
		Save     string     `json:"save"`
	} `json:"search_items"`
}

func loadConfig(path string) (*Config, error) {
	return nil, nil
}
