package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func startRESTServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go server gotry!"))
	})

	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "10000" // Render default
	}
	serverAddr := "0.0.0.0:" + portEnv
	fmt.Printf("Starting REST API server on %s...\n", serverAddr)

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		fmt.Printf("REST API server error: %v\n", err)
	}
}

func main() {
	fmt.Println("\n✓ Connected to WhatsApp try! Type 'help' for commands.")

	// Start the REST server (blocking)
	startRESTServer()

	// Optional: Keep the process alive for graceful shutdown (not needed if server is blocking)
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("REST server is running. Press Ctrl+C to disconnect and exit.")
	<-exitChan
}
