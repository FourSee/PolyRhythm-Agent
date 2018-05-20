package main

import (
	"fmt"
	"os"

	apiclient "github.com/foursee/swagger-go/client"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var exitStatus = 0

var (
	pair                = kingpin.Command("pair", "Pair your device with ShellGame")
	keyBits             = pair.Flag("bitsize", "The RSA key size to generate for the pairing request. Can be either 2048 or 4096. Defaults to 2048").Short('b').Default("2048").Int()
	run                 = kingpin.Command("run", "Run command and recieve notification")
	pipeIn              = kingpin.Command("--", "Read from stdin").Default()
	onStartNotification = run.Flag("onStartNotification", "Send a notifcation on start").Bool()
	command             = run.Arg("Command", "").Required().String()
	commandArgs         = run.Arg("Args", "").Strings()
	polyrhythmAPI       = *apiclient.Default
)

func main() {
	initialize()
	kingpin.Version("0.0.1")

	switch kingpin.Parse() {
	case pair.FullCommand():

		// if config().PairedDevice.PublicKey != "" {
		// 	fmt.Print("Remove existing pairing and create a new one? (Anything other than 'YES' will abort) ")
		// 	if !askForConfirmation() {
		// 		fmt.Println("Aborting")
		// 		return
		// 	}
		// }
		if in_array(*keyBits, []int{2048, 4096}) {
			newPairingRequest()
		} else {
			fmt.Printf("%v is not a permitted keybit size. Permitted sizes are 2048 and 4096\n", *keyBits)
		}

	// Post message
	case run.FullCommand():
		fmt.Printf("%v %v", *commandArgs, *onStartNotification)
		cr := commandRunner{command: *command, args: *commandArgs, startNotification: *onStartNotification}
		cr.run()
	case pipeIn.FullCommand():
		readFromPipe()
	}

}

func initialize() {
	if os.Getenv("POLYRHYTHM_HOST") != "" {
		t := apiclient.DefaultTransportConfig().WithHost(os.Getenv("POLYRHYTHM_HOST"))
		polyrhythmAPI = *apiclient.NewHTTPClientWithConfig(nil, t)
	}
}
