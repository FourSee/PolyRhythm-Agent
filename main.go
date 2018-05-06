package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var start time.Time

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	var (
		cmd *exec.Cmd
		// cmdOut []byte
		err error
	)

	cmd = exec.Command(arg, os.Args[2])

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println("stuff")
	}

	scanner := bufio.NewScanner(stdout)

	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	start = time.Now()

	defer stuff()

	if err = cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}

	fmt.Println("starting!")

	if err := cmd.Wait(); err != nil {
		fmt.Println("Waiting")

		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}
	}

}

func stuff() {
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time Elapsed:", elapsed)
}
