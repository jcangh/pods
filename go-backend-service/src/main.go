package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := fmt.Sprintf("la hora es %v  y hostname es %v", time.Now(), os.Getenv("HOSTNAME"))
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
