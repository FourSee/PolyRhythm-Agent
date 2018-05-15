package main

import (
	qrcode "github.com/Baozisoftware/qrcode-terminal-go"
)

type qrCode struct {
	URL string `json:"url"`
}

func (qrc *qrCode) show() {
	qr := qrcode.New()
	qr.Get(qrc.URL).Print()
}
