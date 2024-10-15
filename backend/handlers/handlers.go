// file: handlers/handlers.go
package handlers

import (
	"compress/gzip"
	"compression/data" // Import the data package
	"encoding/json"
	"net/http"
)

func CreatePayloadHandler(users *[]data.UserInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addCORSHeaders(w)
		json.NewEncoder(w).Encode(users)
	}
}

func CreateCompressedPayloadHandler(users *[]data.UserInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addCORSHeaders(w)

		w.Header().Set("Content-Encoding", "gzip")
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()

		json.NewEncoder(gzipWriter).Encode(users)
	}
}

func addCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
