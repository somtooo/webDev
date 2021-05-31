package main

import (
	"io"
	"net/http"
)

type hotdogg int

func (m hotdogg) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty")
	default:
		io.WriteString(res, "Welcome")
	}
}

func main() {
	var d hotdogg
	http.ListenAndServe(":3000", d)

}
