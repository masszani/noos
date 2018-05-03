package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Frase struttura per input
type Frase struct {
	Testo string `json:"testo,omitempty"`
}

// Nooize - Giulizza la frase
func Nooize(w http.ResponseWriter, req *http.Request) {
	var frase Frase
	_ = json.NewDecoder(req.Body).Decode(&frase)
	fraseNooized := frase.Testo + " ,no?"
	json.NewEncoder(w).Encode(fraseNooized)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/nooize", Nooize).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
