package config

import (
	"encoding/json"
	"os"
)

func Load(fileName string) (config Config, err error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &defaultConfig{data: m}, nil
}
