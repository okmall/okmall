package main

import (
	"github.com/okmall/okmall/cmd/okmall/app"
	"github.com/unisx/logus"
)

func main() {
	defer logus.Sync()
	app.Execute()
}
