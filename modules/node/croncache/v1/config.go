package v1

import (
	"encoding/json"
)

type Config struct {
	LogFilePath string
	TaskManager *json.RawMessage
}

func NewConfig() *Config {
	c := &Config{}
	return c
}
