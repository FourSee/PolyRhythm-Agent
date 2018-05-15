package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var exitStatus = 0

var (
	pair                = kingpin.Command("pair", "Pair your device with ShellGame")
	run                 = kingpin.Command("run", "Run command and recieve notification")
	onStartNotification = run.Flag("onStartNotification", "Send a notifcation on start").Bool()
	command             = run.Arg("Command", "").Required().String()
	commandArgs         = run.Arg("Args", "").Strings()
)

func main() {
	kingpin.Version("0.0.1")

	switch kingpin.Parse() {
	// Register user
	// case qr.FullCommand():
	// 	content := "http://localhost:3000/v1/pr/4nnT2ItM5QSDn2Ziu78jit"
	// 	obj := qrcode.New2(qrcode.ConsoleColors.NormalBlack, qrcode.ConsoleColors.BrightWhite, qrcode.QRCodeRecoveryLevels.Low)
	// 	obj.Get([]byte(content)).Print()
	case pair.FullCommand():
		newPairingRequest()

	// Post message
	case run.FullCommand():
		fmt.Printf("%v %v", *commandArgs, *onStartNotification)
		cr := commandRunner{command: *command, args: *commandArgs, startNotification: *onStartNotification}
		cr.run()
	}
}
