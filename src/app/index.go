package main

import (
	"../server"
	"../driver"
)

func main() {
	driver.MakeConnections()
	server.HandleRequests()
}
