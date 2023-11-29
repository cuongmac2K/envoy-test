package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", auth)
	http.ListenAndServe(":5000", nil)
}

func auth(writer http.ResponseWriter, request *http.Request) {
	token := request.FormValue("Token")
	if token == "abc" {
		writer.Write([]byte("ok r nha"))
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(403)
	}
}
