package main

import (
	"github.com/okmall/okmall/cmd/core"
	"github.com/unisx/logus"

	_ "github.com/okmall/okmall/cmd/backup"
	_ "github.com/okmall/okmall/cmd/serve"
	_ "github.com/okmall/okmall/cmd/version"
)

func main() {
	defer logus.Sync()

	// Setup root cli command of application
	core.Setup(
		"okmall",                       // command name
		"Provide okmall agent service", // command short describe
		"Provide okmall agent service", // command long describe
	)

	// Execute start application
	core.Execute()
}
