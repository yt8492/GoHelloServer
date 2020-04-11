package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RootResponse struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		res := RootResponse {
			Text: "hello, world",
		}
		json.NewEncoder(w).Encode(res)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}