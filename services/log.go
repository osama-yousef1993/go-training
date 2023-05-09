package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type logger struct {
	Severity string            `json:"severity"`
	Message  string            `json:"message"`
	Labels   map[string]string `json:"labels"`
}

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func (l logger) String() string {
	l.Message = strings.ReplaceAll(l.Message, "\"", "")
	if len(l.Labels) == 0 {
		return fmt.Sprintf("{\"severity\": \"%v\",\"message\": \"%v\"}", l.Severity, l.Message)
	} else {
		lb, err := json.Marshal(l.Labels)
		if err != nil {
			return fmt.Sprintf("{\"severity\": \"%v\",\"message\": \"%v\"}", l.Severity, l.Message)
		}
		return fmt.Sprintf("{\"severity\": \"%v\",\"message\": \"%v\", \"labels\": %v}", l.Severity, l.Message, string(lb))
	}
}

func Info(body string, t ...interface{}) {
	log.Printf(logger{Severity: "INFO", Message: fmt.Sprintf(body, t...)}.String())
}
