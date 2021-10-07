package main

import (
	"os"

	"github.com/schaermu/gopress/cmd"
)

var osExit = os.Exit

func main() {
	osExit(cmd.EchoStdErrIfError(cmd.Execute()))
}
