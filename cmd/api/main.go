package main

import (
	"net/http"

	"github.com/Taofeekabdulazeez/anagramsolver/internal/api"
)

func main() {
	http.HandleFunc("/{word}", api.Handler)

	http.ListenAndServe(":8080", nil)
}
