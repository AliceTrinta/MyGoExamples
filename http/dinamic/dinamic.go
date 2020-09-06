package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hour(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1> Hour: %s</h1>", s)
}

func main() {
	http.HandleFunc("/Hour", hour)
	log.Println("Executing...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
