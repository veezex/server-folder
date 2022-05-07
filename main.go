package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "server port")
	flag.Parse()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
