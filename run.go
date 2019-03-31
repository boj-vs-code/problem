package main

import (
	"github.com/moreal/boj-vs-code-api-server/server"
	"google.golang.org/appengine"
)

func main() {
	s := server.CreateServer()
	s.Run(":8080")
	appengine.Main()
}