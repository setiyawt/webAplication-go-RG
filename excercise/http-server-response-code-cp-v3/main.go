package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(Writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			Writer.WriteHeader(http.StatusMethodNotAllowed)
			Writer.Write([]byte("Method is not allowed"))
			return
		}
		name := request.URL.Query().Get("name")
		if name == "" {
			Writer.WriteHeader(http.StatusNotFound)
			Writer.Write([]byte("Data not found"))
			return
		}
		exist := IsNameExists(name)
		if !exist {
			Writer.WriteHeader(http.StatusNotFound)
			Writer.Write([]byte("Data not found"))
			return
		}
		Writer.WriteHeader(http.StatusOK)
		Writer.Write([]byte("Name is exists"))
	} // TODO: replace this
}

func Hello() http.HandlerFunc {
	return func(Writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("name")
		if name == "" {
			Writer.WriteHeader(http.StatusNotFound)
			Writer.Write([]byte("Data not found"))
			return
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/students", CheckStudentName())

	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
