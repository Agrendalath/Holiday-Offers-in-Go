package main

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
)

// Handle endpoint.
func handler(w http.ResponseWriter, r *http.Request) {
	data, err := filterData(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Set("Content-Type", "application/json")

	var enc *json.Encoder
	if r.URL.Query().Get("debug") != "true" {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		enc = json.NewEncoder(gz)
		defer gz.Close()
	} else {
		enc = json.NewEncoder(w)
		enc.SetIndent("", "    ")
	}

	enc.Encode(data)
}

func main() {
	http.HandleFunc("/v1/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
