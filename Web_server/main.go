package main

import (
	"log"
	"os"
)

func main() {
	server := NewServer(":3000")
	server.Handle("POST", "/mutant/", server.AddMiddleware(isMutant))
	server.Handle("GET", "/stats/", server.AddMiddleware(Stats))
	server.Listen()
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
