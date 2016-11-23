package main

import (
	"challenge2/rodo-training-2/ipScanner"
	"net/http"
)

func main() {
	errorHandler := ipScanner.ErrorHandler{}
	http.ListenAndServe(":12345", errorHandler)
}