package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/VicentiuCristianBadea/go-app.git/details"
	"github.com/gorilla/mux"
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check")
	response := map[string]string{
		"status": "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

// Root handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

// Details handler
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, errHostName := GetHostName()
	if errHostName != nil {
		panic(errHostName)
	}
	IP, errIP := GetOutboundIP()
	if errIP != nil {
		panic(errIP)
	}

	response := map[string]string{
		"hostname": hostname,
		"ip": IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

// Main function
func main() {
    r := mux.NewRouter()

    r.HandleFunc("/health", healthHandler) 
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Println("Starting the server on port 80")

	log.Fatal(http.ListenAndServe(":80", r))
}