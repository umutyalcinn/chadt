package main

import (
	"github.com/umutyalcinn/chadt/server"
)



func main() {
	server := server.NewServer(":6969")

	server.Start()
}

