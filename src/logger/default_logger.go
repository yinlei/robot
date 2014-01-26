package logger

import (
	"log"
)

type DefaultLogger struct {
}

func NewDefaultLogger() Logger {
	return &DefaultLogger{}
}

func (dl *DefaultLogger) Debug(info, message string) {
	log.Println("[D]", info, message)
}

func (dl *DefaultLogger) Info(info, message string) {
	log.Println("[I]", info, message)
}

func (dl *DefaultLogger) Warning(info, message string) {
	log.Println("[W]", info, message)
}

func (dl *DefaultLogger) Error(info, message string) {
	log.Println("[E]", info, message)
}

func (dl *DefaultLogger) Critical(info, message string) {
	log.Println("[C]", info, message)
}
