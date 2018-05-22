package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/djherbis/stream"
)

func readFromPipe(silent bool) {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: some_process | polyrhythm ")
		return
	}
	buf := bufio.NewReader(os.Stdin)
	enc, stdOut := duplicateReader(buf)

	if !silent {
		go io.Copy(os.Stdout, stdOut)
	}

	send(enc, 0, "")
}

func duplicateReader(r io.Reader) (r1, r2 io.Reader) {
	buf := make([]byte, 0, 4*1024)
	w, err := stream.New("mystream")
	check(err)

	r1, err = w.NextReader()
	check(err)

	r2, err = w.NextReader()
	check(err)

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		w.Write(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	defer w.Close()
	return
}
