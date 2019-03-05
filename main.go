package main

import (
	"fmt"

	openapiGo "github.com/foursee/openapiGo"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var exitStatus = 0

var (
	pair                = kingpin.Command("pair", "Pair your device with ShellGame")
	quiet               = kingpin.Flag("quiet", "Silences stdout").Short('q').Default("false").Bool()
	run                 = kingpin.Command("run", "Run command and recieve notification")
	pipeIn              = kingpin.Command("--", "Read from stdin").Default()
	onStartNotification = run.Flag("onStartNotification", "Send a notifcation on start").Bool()
	command             = run.Arg("Command", "").Required().String()
	commandArgs         = run.Arg("Args", "").Strings()
	polyrhythmAPI       = openapiGo.NewAPIClient(openapiGo.NewConfiguration())
)

func main() {
	initialize()
	kingpin.Version("0.0.1")

	switch kingpin.Parse() {
	case pair.FullCommand():
		if config().PairedDevice.PublicKey != "" {
			fmt.Print("Remove existing pairing and create a new one? (Anything other than 'YES' will abort) ")
			if !askForConfirmation() {
				fmt.Println("Aborting")
				return
			}
		}
		newPairingRequest()
	case run.FullCommand():
		fmt.Printf("%v %v", *commandArgs, *onStartNotification)
		cr := commandRunner{command: *command, args: *commandArgs, startNotification: *onStartNotification}
		cr.run()
	case pipeIn.FullCommand():
		readFromPipe(*quiet)
	}
}

func initialize() {
	// if os.Getenv("POLYRHYTHM_HOST") != "" {
	// 	t := apiclient.DefaultTransportConfig().WithHost(os.Getenv("POLYRHYTHM_HOST"))
	// 	polyrhythmAPI = *apiclient.NewHTTPClientWithConfig(nil, t)
	// }
}
