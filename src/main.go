package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	appName string = "N/A"
	version string = "N/A"
	sha     string = "N/A"
)

type Info struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
	SHA     string `json:"sha"`
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := Info{
		AppName: appName,
		Version: version,
		SHA:     sha,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(info); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
}
