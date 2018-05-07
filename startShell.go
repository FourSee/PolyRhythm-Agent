package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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
	b, err := json.Marshal(ss)

	if err != nil {
		log.Fatal("Failed to marshal StartShell")
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
