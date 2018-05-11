package main

import (
	"encoding/json"
	"log"
	"net/http"

	qrcode "github.com/Baozisoftware/qrcode-terminal-go"
)

type qrCode struct {
	URL string `json:"url"`
}

func (qrc *qrCode) show() {
	qr := qrcode.New()
	qr.Get(qrc.URL).Print()
}

func getPairingURL() qrCode {
	resp, err := http.Get(apiURL)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	d := json.NewDecoder(resp.Body)

	var qrc qrCode
	err = d.Decode(&qrc)

	if err != nil {
		log.Fatalf("Failed to decode data %v", err.Error())
	}

	return qrc
}
