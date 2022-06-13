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
	//Decode request
	var body = DecoderBody(r)
	//Body validations
	if mutant.ValidateMatrizDimensions(body.Dna) {
		//Valid mutant
		var isMutant bool = mutant.IsMutant(body.Dna)
		body.IsMutant = isMutant

		database.SaveTransactions(body)

		if isMutant {
			Response(w, 200)
		} else {
			Response(w, 403)
		}
	} else {
		Response(w, 400)
	}
}

func Response(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func DecoderBody(r *http.Request) mutant.MutantInfo {
	var dna mutant.MutantInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dna)
	if err != nil {
		panic(err)
	}
	return dna
}
