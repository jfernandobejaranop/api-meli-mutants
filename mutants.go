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
	var req Mutant

	err := decoder.Decode(&req)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	//saveTransactions(req.Dna)

	if IsMutant(req.Dna) {
		Response(w, 200)
	} else {
		Response(w, 403)
	}
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
	Dna []string `json:"dna"`
}
