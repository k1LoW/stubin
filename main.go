package main

import (
	"os"

	"github.com/k1LoW/stubin/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args, os.Stdout, os.Stderr))
}
