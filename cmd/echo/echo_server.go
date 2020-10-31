package main

import (
	"log"
	"mock-grpc/cmd/echo/app"
)

func main() {
	command := app.NewEchoServerCommand()
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
