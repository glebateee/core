package config

import (
	"encoding/json"
	"os"
	"strings"
)

func Load(fileName string) (config Config, err error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.NewDecoder(strings.NewReader(string(data))).Decode(&m); err != nil {
		return nil, err
	}
	return &defaultConfig{data: m}, nil
}
