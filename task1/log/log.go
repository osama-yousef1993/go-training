package log

import (
	"log"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func Err(err string) {
	now := time.Now()
	log.Fatalf("error: (%s) %s\n", now.Format(timeFormat), err)
}

func Info(message string) {
	now := time.Now()
	log.Printf("info: (%s) %s\n", now.Format(timeFormat), message)
}
