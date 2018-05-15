package main

import (
	"encoding/json"
	"log"
	"time"
)

// EndShell is struct that gets sent at the end of the running the application
type EndShell struct {
	EndDate    time.Time
	Pid        int
	Elapsed    time.Duration
	Output     string
	ExitStatus int
}

func (es *EndShell) send() {
	_, err := json.Marshal(es)

	if err != nil {
		log.Fatal("Failed to marshal StartShell")
	}

	// req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
}

func (es *EndShell) encodeData() {

}

func (es *EndShell) setElapsed(sd time.Time, ed time.Time) {
	es.Elapsed = ed.Sub(sd)
	es.EndDate = ed
}
