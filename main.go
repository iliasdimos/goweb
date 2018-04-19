//go:generate stringer -type=Color
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var version = "v2"

const (
	_ Color = iota
	red
	blue
	yellow
	pink
	white
)

type Color int

var color Color

func main() {
	port := ":" + getPort()
	color = Color(rand.Intn(5))

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
	fmt.Fprintf(w, "<!DOCTYPE html><html><h2>Simple Go Web Server %s</h2><body style='background-color: %s'> Running on %s</body></html>", version, color.String(), host)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
