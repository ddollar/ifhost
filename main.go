package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func check(err error) {
	if err != nil {
		die(err)
	}
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

func usage() {
	die(fmt.Errorf("usage: ifhost <hostname> <command...>"))
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	host, err := os.Hostname()
	check(err)

	name := strings.Split(host, ".")[0]

	if name == os.Args[1] {
		cmd := exec.Command(os.Args[2], os.Args[3:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
