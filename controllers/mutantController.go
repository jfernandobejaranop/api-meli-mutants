package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-meli-mutants/database"
	"github.com/api-meli-mutants/mutant"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "technical challenge for MELI")
}
func Statistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(database.StatisticsBD())
}

func Mutants(w http.ResponseWriter, r *http.Request) {
	var body = DecoderBody(r)
	var isMutant bool = mutant.IsMutant(body.Dna)

	body.IsMutant = isMutant
	database.SaveTransactions(body)

	if isMutant {
		Response(w, 200)
	} else {
		Response(w, 403)
	}
}

func Response(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func DecoderBody(r *http.Request) mutant.Mutant {
	var dna mutant.Mutant
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dna)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	return dna
}
