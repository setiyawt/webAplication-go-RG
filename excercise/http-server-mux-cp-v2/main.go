package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("Monday, 2 January 2006")
		response := fmt.Sprintf(currentTime)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))

	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("Hello there"))
		} else {
			response := fmt.Sprintf("Hello, %s!", name)
			w.Write([]byte(response))
		}
		w.WriteHeader(http.StatusOK)
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/time", TimeHandler())
	mux.Handle("/hello", SayHelloHandler())
	// TODO: answer here

	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
