package main

import (
	"compression/data"
	"compression/handlers"
	"fmt"
	"net/http"
)

func main() {
	_100users95CompressionRatio := data.GeneratePayload(100, 95)
	http.HandleFunc("/payload", handlers.CreatePayloadHandler(&_100users95CompressionRatio))
	http.HandleFunc("/payloadC", handlers.CreateCompressedPayloadHandler(&_100users95CompressionRatio))

	_1000users95CompressionRatio := data.GeneratePayload(1000, 95)
	http.HandleFunc("/payload2", handlers.CreatePayloadHandler(&_1000users95CompressionRatio))
	http.HandleFunc("/payload2C", handlers.CreateCompressedPayloadHandler(&_1000users95CompressionRatio))

	_10000users95CompressionRatio := data.GeneratePayload(10_000, 95)
	http.HandleFunc("/payload3", handlers.CreatePayloadHandler(&_10000users95CompressionRatio))
	http.HandleFunc("/payload3C", handlers.CreateCompressedPayloadHandler(&_10000users95CompressionRatio))

	fmt.Println("Started 8080")
	http.ListenAndServe(":8080", nil)
}
