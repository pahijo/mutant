package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	println("Initialize Logger")
	f, err := os.OpenFile(fmt.Sprintf("%v.log", os.Args[0]), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	if port == "" {
		//log.Fatal("PORT must be set")
		port = "3000"
	}
	args := strings.Join(os.Args[1:], ", ")
	log.Println("App started, Port:", port, "Args:", args)

	server := NewServer(":" + port)
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
