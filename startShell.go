package main

import (
	"encoding/json"
	"log"
	"time"
)

// StartShell is struct that gets sent at the beginning of the running the application
type StartShell struct {
	StartDate        time.Time `json:"startDate"`
	Pid              int       `json:"pid"`
	SendNotification bool      `json:"sendNotification"`
	Title            string    `json:"title"`
}

func (ss *StartShell) send() {
	_, err := json.Marshal(ss)

	if err != nil {
		log.Fatal("Failed to marshal StartShell")
	}
}
