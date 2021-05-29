package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	reset = "\033[0m"
	bold = "\033[1m"
)

func isRoot() bool {
  return os.Getuid() == 0
}

func checkRoot() {
  if !isRoot() {
    err("Shade must be ran as root")
    os.Exit(1)
  }
}

func rev(s string) string {
  return "\033[7m" + s + "\033[0m";
}

func success(s ...string)  {
  fmt.Printf("%s\033[32m%s%s\n", bold, strings.Join(s, " "), reset)
}

func info(s ...string) {
  fmt.Printf("%s\033[34m%s%s\n", bold, strings.Join(s, " "), reset)
}

func err(s ...string)  {
  fmt.Printf("%s\033[31m%s%s\n", bold, strings.Join(s, " "), reset)
}

func run(cmdstr string) {
	cmdArgs := strings.Fields(cmdstr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	fmt.Printf("Command finished with error: %v\n", err)
	return
}

