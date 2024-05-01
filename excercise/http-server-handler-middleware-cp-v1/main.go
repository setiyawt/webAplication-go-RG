package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
		return
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
		}

		return
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	http.HandleFunc("/student", StudentHandler())
	http.ListenAndServe("localhost:8080", RequestMethodGet(nil))
	http.ListenAndServe("localhost:8080", nil)

}
