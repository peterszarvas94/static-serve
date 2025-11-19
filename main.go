package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var (
		dir  = flag.String("d", ".", "Directory to serve")
		port = flag.String("p", "8080", "Port to listen on")
	)

	// Add long flag aliases
	flag.StringVar(dir, "dir", ".", "Directory to serve")
	flag.StringVar(port, "port", "8080", "Port to listen on")

	flag.Parse()

	// Convert to absolute path for better logging
	absDir, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatalf("Error resolving directory path: %v", err)
	}

	// Check if directory exists
	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", absDir)
	}

	// Create file server
	fs := http.FileServer(http.Dir(absDir))
	http.Handle("/", fs)

	fmt.Printf("Serving %s on http://localhost:%s\n", absDir, *port)
	fmt.Printf("Press Ctrl+C to stop the server\n")

	// Start server
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
