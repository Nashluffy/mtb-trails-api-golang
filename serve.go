package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type metadata struct {
	Timestamp string `json:"query_time"`
}

type response struct {
	Metadata metadata `json:"metadata"`
	Data []Trail  `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s connected with %s", r.RemoteAddr, r.Header.Get("User-Agent"))

	var (
		trails []Trail
		responseMetadata metadata
		finalResponse response
	)

	trails = FetchTrailStatus()

	responseMetadata = metadata {
		Timestamp: time.Now().Format("2006.01.02 15:04:05"),
	}

	finalResponse = response {
		Metadata: responseMetadata,
		Data:     trails,
	}

	marshaledResponse, err := json.MarshalIndent(finalResponse, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledResponse)
}

func main() {
	http.HandleFunc("/trails", handler)
	fmt.Println("Serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
