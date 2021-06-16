package main

import (
	"cyoawebapp/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	entryPoint := flag.String("start", "intro", "the entry point to the CYOA title")
	port := flag.Int("port", 8080, "the port to start CYOA web application on")
	fileName := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, *entryPoint)
	fmt.Printf("Starting the server on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
