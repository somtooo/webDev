package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "dog dog")
	})

	http.HandleFunc("/cat", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "cat cat cat")
	})
	http.HandleFunc("/", dog)

	http.ListenAndServe(":3000", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}
