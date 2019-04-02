package main

import (
	"github.com/moreal/boj-vs-code-api-server/server"
)

func main() {
	s := server.CreateServer()
	s.Run(":8080")
}