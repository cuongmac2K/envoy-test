package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/get", getUser)
	http.ListenAndServe(":4000", nil)
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("ok r nha"))
	writer.WriteHeader(200)
}
