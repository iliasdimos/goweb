package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var version = "v1"

func main() {
	port := ":" + getPort()

	// Handling signals
	signals := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go catchSignal(signals, done)
	// End of handling signals

	http.HandleFunc("/", indexHandler)
	fmt.Printf("Starting server at %s\n", port)
	fmt.Println("Press ctrl-c to terminate...")
	http.ListenAndServe(port, nil)

	<-done
	fmt.Println("Closing now ! Bye !")
	os.Exit(1)
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

func catchSignal(ch chan os.Signal, done chan bool) {
	sig := <-ch

	fmt.Println("sig received:", sig)

	switch sig {
	case syscall.SIGINT:
		fmt.Println("Handling a SIGINT here")
	case syscall.SIGTERM:
		fmt.Println("Handling a SIGTERM here")
	default:
		fmt.Println("Other signal received")
	}

	done <- true
}
