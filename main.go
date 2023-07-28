package main

import (
	"os"

	"github.com/linuxsuren/github-go/cmd"
)

func main() {
	if err := cmd.NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
