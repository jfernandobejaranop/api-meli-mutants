package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "technical challenge for MELI")
}
func Statistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Stats())
}

func mutant(w http.ResponseWriter, r *http.Request) {
	var body = DecoderBody(r)
	var isMutant bool = IsMutant(body.Dna)

	body.IsMutant = isMutant
	SaveTransactions(body)

	if isMutant {
		Response(w, 200)
	} else {
		Response(w, 403)
	}
}

func DecoderBody(r *http.Request) Mutant {
	var dna Mutant
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dna)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	return dna
}

func IsMutant(dna []string) bool {
	if HorizontalArray(dna) {
		fmt.Println("entraHorizontal")
		return true
	} else if VerticalArray(dna) {
		fmt.Println("entraVertical")
		return true
	} else if RightDiagonalArray(dna) {
		return true
	} else if LeftDiagonalArray(dna) {
		return true
	}
	return false
}

func Response(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

type Mutant struct {
	Dna      []string `json:"dna"`
	IsMutant bool
}
