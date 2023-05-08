package log

import (
	"log"
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func LogErr(msg string) {
	now := time.Now()
	log.Fatalf("error: (%s) %s\n", now.Format(DDMMYYYYhhmmss), msg)
}

func LogInfo(msg string) {
	now := time.Now()
	log.Printf("info: (%s) %s\n", now.Format(DDMMYYYYhhmmss), msg)
}
