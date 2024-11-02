package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RequestDetails struct {
	Method     string              `json:"method"`
	URL        string              `json:"url"`
	Headers    map[string][]string `json:"headers"`
	RemoteAddr string              `json:"remote_addr"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("new request on:", r.URL.String())
	details := RequestDetails{
		Method:     r.Method,
		URL:        r.URL.String(),
		Headers:    r.Header,
		RemoteAddr: r.RemoteAddr,
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "8080")
	addr := fmt.Sprintf(":%s", port)
	http.HandleFunc("/", handler)
	log.Printf("starting go-inspect on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
