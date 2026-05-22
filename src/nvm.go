package main

import (
	"fmt"
	"os"

	"github.com/coreybutler/nvm-windows/src/cli"
	"github.com/coreybutler/nvm-windows/src/nvm"
)

const (
	// Version is the current version of nvm-windows
	Version = "1.1.12"
)

func main() {
	// Initialize the NVM environment
	env := nvm.NewEnvironment()
	if err := env.Setup(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing nvm: %v\n", err)
		os.Exit(1)
	}

	// Parse and execute the CLI command
	app := cli.NewApp(Version, env)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
