package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_qrCode_getPairingURL(t *testing.T) {
	testURL := "http://someplace.com"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected ‘GET’ request, got ‘%s’", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"url": "%s"}`, testURL)
	}))
	defer ts.Close()

	apiURL = ts.URL

	qrT := getPairingURL()
	if qrT.URL != testURL {
		t.Errorf("Raw Value %v", qrT)
		t.Errorf("Expected urls to be equal %v, got ‘%v’", testURL, qrT)
	}
}
