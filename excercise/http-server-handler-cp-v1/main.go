package main

import (
	"net/http"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Tuesday, 30 April 2024"))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
