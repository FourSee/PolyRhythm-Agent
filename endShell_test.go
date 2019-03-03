package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_EndShell_SendData(t *testing.T) {
	now := time.Now().UTC()
	es := EndShell{Elapsed: 20, EndDate: now, ExitStatus: 0, Output: "", Pid: 100}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}

		t.Log(r.Body)

		d := json.NewDecoder(r.Body)

		var esT EndShell
		err := d.Decode(&esT)
		if err != nil {
			t.Log("Failed to decode")
			panic(err)
		}

		defer r.Body.Close()

		if esT != es {
			t.Errorf("Expected structs to be equal %v, got ‘%v’", es, esT)
		}

		t.Log(esT)

		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	es.send()
}
