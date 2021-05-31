package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("blah")
}

func main() {
	var d hotdog
	http.ListenAndServe(":3000", d)
}
