package mutant

func IsMutant(dna []string) bool {
	if HorizontalArray(dna) {
		return true
	} else if VerticalArray(dna) {
		return true
	} else if RightDiagonalArray(dna) {
		return true
	} else if LeftDiagonalArray(dna) {
		return true
	}
	return false
}
