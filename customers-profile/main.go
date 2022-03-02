package main

import (
	"github.com/custs-risk/connection"
	"github.com/custs-risk/handlers"
)

func main() {
	connection.Connect()
	handlers.HandleReq()
}
