package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor GO")
}

func mutant(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dna Mutant

	err := decoder.Decode(&dna)

	if err != nil {
		panic(err)
	}

	saveTransactions(dna)

	defer r.Body.Close()

	if ismutant(dna) {
		Response(w, 200)
	} else {
		Response(w, 403)
	}
}

func Response(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func ismutant(req Mutant) bool {
	return false
}

type Mutant struct {
	Dna []string `json:"dna"`
}
