package main

import (
	"github.com/alimy/logus"
	"github.com/okmall/okmall/cmd/okmall/app"
)

func main() {
	defer logus.Sync()
	app.Execute()
}
