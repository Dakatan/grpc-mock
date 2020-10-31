package main

import (
	"log"
	"mock-grpc/cmd/session/app"
)

func main() {
	command := app.NewSessionServerCommand()
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
