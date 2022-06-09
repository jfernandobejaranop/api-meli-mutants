package main

func HorizontalArray(dna []string) bool {
	return matchValidate(dna)
}

func VerticalArray(dna []string) bool {
	var row string
	var arrB []string

	for fila := 0; fila < len(dna); fila++ {
		row = ""

		for letra := 0; letra < len(dna); letra++ {
			row = row + dna[letra][fila:fila+1] //matriz 2
		}
		arrB = append(arrB, row)
	}
	return matchValidate(arrB)
}
