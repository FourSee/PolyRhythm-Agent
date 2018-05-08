package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var apiURL = "http://google.ca"
var exitStatus = 0

var (
	debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
	timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
	ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
	count   = kingpin.Arg("count", "Number of packets to send").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]

	// arg := os.Args[1]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

	// cr := commandRunner{command: arg, args: os.Args[2]}
	// cr.run()
}
