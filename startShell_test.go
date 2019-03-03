package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_StartShell_SendData(t *testing.T) {
	now := time.Now().UTC()
	ss := StartShell{Pid: 100, SendNotification: false, Title: "Cheese", StartDate: now}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}

		t.Log(r.Body)

		d := json.NewDecoder(r.Body)

		var ssT StartShell
		err := d.Decode(&ssT)
		if err != nil {
			t.Log("Failed to decode")
			panic(err)
		}

		defer r.Body.Close()

		if ssT != ss {
			t.Errorf("Expected structs to be equal %v, got ‘%v’", ss, ssT)
		}

		t.Log(ssT)

		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	ss.send()
}
