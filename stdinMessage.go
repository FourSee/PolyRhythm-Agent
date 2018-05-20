package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFromPipe() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}
	reader := bufio.NewReader(os.Stdin)

	send(reader, 0, "")

	// for {
	// 	input, _, err := reader.ReadRune()
	// 	if err != nil && err == io.EOF {
	// 		fmt.Printf("FINISHED WITH ENV VARS\n%v\n", strings.Join(os.Environ()[:], "\n\n"))
	// 		break
	// 	}
	// 	fmt.Print(string(input))
	// }
}
