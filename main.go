package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/{word}", Handler)

	http.ListenAndServe(":8080", nil)
}
