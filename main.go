package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Frase struttura per input
type Frase struct {
	Testo string `json:"testo,omitempty"`
	Tipo  string `json:"tipo,omitempty"`
}

func NDecorator(testo string) string {
	return testo + " ,No?"
}

func EDecorator(testo string) string {
	return testo + " ,Eh?"
}

func GDecorator(testo string) string {
	return testo + " ,Giusto?"
}

// Nooize - Giulizza la frase
func Nooize(w http.ResponseWriter, req *http.Request) {
	var frase Frase
	_ = json.NewDecoder(req.Body).Decode(&frase)
	fraseNooized := ""
	switch frase.Tipo {
	case "N":
		fraseNooized = NDecorator(frase.Testo)
	case "E":
		fraseNooized = EDecorator(frase.Testo)
	case "G":
		fraseNooized = GDecorator(frase.Testo)
	}
	json.NewEncoder(w).Encode(fraseNooized)
}

// AutoG - Let pc Giulizze for you
func AutoG(w http.ResponseWriter, req *http.Request) {
	var frase Frase
	_ = json.NewDecoder(req.Body).Decode(&frase)
	fraseNooized := ""
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3)
	switch num {
	case 0:
		fraseNooized = NDecorator(frase.Testo)
	case 1:
		fraseNooized = GDecorator(frase.Testo)
	default:
		fraseNooized = EDecorator(frase.Testo)
	}
	json.NewEncoder(w).Encode(fraseNooized)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/nooize", Nooize).Methods("POST")
	router.HandleFunc("/autog", AutoG).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
