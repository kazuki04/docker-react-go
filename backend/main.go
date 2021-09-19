package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8000", nil)
}
