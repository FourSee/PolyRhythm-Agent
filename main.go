package main

import (
	"fmt"
	"os"
)

var apiURL = "http://google.ca"
var exitStatus = 0

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	cr := commandRunner{command: arg, args: os.Args[2]}
	cr.run()
}
