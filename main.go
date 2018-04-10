package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "v1"

func main() {
	port := ":" + getPort()

	http.HandleFunc("/", indexHandler)
	fmt.Printf("Starting server at %s\n", port)
	fmt.Println("Press ctrl-c to terminate...")
	log.Fatal(http.ListenAndServe(port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "<!DOCTYPE html><html><h2>Simple Go Web Server %s</h2><body style='background-color: %s'> Running on %s</body></html>", version, getColor(), host)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func getColor() string {
	color := os.Getenv("COLOR")
	if color == "" {
		return "white"
	}
	return color
}
