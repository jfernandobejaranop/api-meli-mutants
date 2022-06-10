package mutants

import (
	"fmt"
)

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

type Mutant struct {
	Dna      []string `json:"dna"`
	IsMutant bool
}
