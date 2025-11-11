package main

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	word := r.PathValue("word")

	encoder := json.NewEncoder(w)

	if !IsValidWordParam(word) {
		encoder.Encode(&ErrorResponse{
			Message: "Invalid Request!",
		})
		return
	}

	anagrams := FindAnagrams(word)

	encoder.Encode(&SuccessResponse{
		Words:   anagrams,
		Results: len(anagrams),
	})
}
