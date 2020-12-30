package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"trails-api/scrape"
)


func handler(w http.ResponseWriter, r *http.Request) {
	marshaledTrails, _ := json.Marshal(scrape.FetchTrails())
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledTrails)
}

func main() {
	http.HandleFunc("/trails", handler)
	fmt.Println("Serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}