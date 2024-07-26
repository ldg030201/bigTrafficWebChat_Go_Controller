package main

import (
	"chat_controller_server/cmd/app"
	"chat_controller_server/config"
	"flag"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	a := app.NewApp(c)
	a.Start()
}
