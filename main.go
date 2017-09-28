package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":" + getPort()

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(port, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "<!DOCTYPE html><html><h2>Simple Go Web Server</h2><body style='background-color: %s'> Running on %s</body></html>", getColor(), host)

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
