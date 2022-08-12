package main

import "github.com/VadimGossip/tcpBsonServerExample/internal/app"

var configDir = "config"

func main() {
	app.Run(configDir)
}
