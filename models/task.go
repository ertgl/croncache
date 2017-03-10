package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Name                   string
	Command                string
	Args                   []string
	Interval               Duration
	Timeout                Duration
	IterationOnFail        int
	LogFilePath            string
	CacheEngineModuleName  string
	CacheEngineCredentials *json.RawMessage
	logger                 *log.Logger
}

func (t *Task) Initialize() error {
	f, err := os.OpenFile(t.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 755)
	if err != nil {
		return err
	}
	t.logger = log.New(f, fmt.Sprintf("[%s] ", t.Name), log.LstdFlags|log.Llongfile)
	t.logger.SetOutput(f)
	return err
}

func (t *Task) Log() *log.Logger {
	return t.logger
}
