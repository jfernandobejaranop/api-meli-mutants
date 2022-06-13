package mutant

import (
	"fmt"
	"strings"
)

var letra string
var arrC [][]string

const t = 1000

func HorizontalArray(dna []string) bool {
	fmt.Println("horizontal ", dna)
	return MatchValidateDNA(dna)
}

func VerticalArray(dna []string) bool {
	var row string
	var arrVer []string

	for fila := 0; fila < len(dna); fila++ {
		row = ""
		for letra := 0; letra < len(dna); letra++ {
			row = row + dna[letra][fila:fila+1]
		}
		arrVer = append(arrVer, row)
	}
	fmt.Println("Vertical ", arrVer)
	return MatchValidateDNA(arrVer)
}

func RightDiagonalArray(dna []string) bool {

	var arrMD [t][t]string //Matriz de Diagonales Derechas
	var arrD1 []string     //Array Diagonal Part 1
	var arrD2 []string     //Array Diagonal Part 2
	lenDna := len(dna)

	arrC = nil
	//Lista a Matriz
	for fil := 0; fil < lenDna; fil++ {
		arrC = append(arrC, strings.Split(dna[fil], ""))
	}

	//Primera parte matriz
	for fil := 0; fil < lenDna; fil++ {
		for col := lenDna - 1; col >= 0; col-- {
			letra = letra + arrC[fil][col]
			if fil >= col {
				break
			}
		}
		arrD1 = append(arrD1, letra)
		letra = ""
	}

	//Segunda parte matriz
	for fil := 0; fil < lenDna; fil++ {
		for col := 1; col < lenDna; col++ {
			if (fil + col) >= lenDna {
				break
			}
			letra = letra + arrC[fil+col][fil]
		}
		arrD1[fil] = arrD1[fil] + letra
		letra = ""
	}

	//Pivoteo matriz
	for fil := 0; fil < len(arrD1); fil++ {
		splits := strings.Split(arrD1[fil], "")
		for col := range splits {
			if !(splits[col] == "") {
				arrMD[col+fil][fil] = splits[col]
			}
		}
	}

	// Pivote final
	for fil := range arrMD {
		for col := 0; col < len(arrMD[fil]); col++ {
			letra = letra + arrMD[fil][col]
		}
		arrD2 = append(arrD2, letra)
		letra = ""
	}

	fmt.Println("iagonalDerechas", arrD2)
	return MatchValidateDNA(arrD2)
}

func LeftDiagonalArray(dna []string) bool {
	lenDna := len(dna)
	var arrML [t][t]string // Matriz Diagonales Izquierdas
	var arrL1 []string     // Array Diagonal Part 1
	var arrL2 []string     // Array Diagonal Part 2
	var cont int = 0

	//PRIMERA PARTE MATRIZ
	for fil := range arrC {
		for col := 0; col < len(arrC[fil]); col++ {
			if col >= len(arrC[fil])-fil {
				break
			}
			letra = letra + arrC[col][fil]
		}
		arrL1 = append(arrL1, letra)
		letra = ""
	}

	//SEGUNDA PARTE MATRIZ
	for fil := lenDna - 1; fil >= 1; fil-- {
		for col := 1; col < lenDna; col++ {
			if col+cont > lenDna-1 {
				break
			}
			letra = letra + arrC[col+cont][fil]
		}
		arrL1[cont] = arrL1[cont] + letra
		letra = ""
		cont++
	}

	//PIVOTEO MATRIZ COMPLETA
	for fil := 0; fil < len(arrL1); fil++ {
		splits := strings.Split(arrL1[fil], "")
		for col := range splits {
			if !(splits[col] == "") {
				arrML[col+fil][fil] = splits[col]
			}
		}
	}

	// PIVOTE FINAL
	for fil := range arrML {
		for col := 0; col < len(arrML[fil]); col++ {
			letra = letra + arrML[fil][col]
		}
		arrL2 = append(arrL2, letra)
		letra = ""
	}

	fmt.Println("Diagonal izquierda", arrL2)
	return MatchValidateDNA(arrL2)
}
func ValidateMatrizDimensions(dna []string) bool {
	for i := 0; i < len(dna); i++ {
		if len(dna[i]) != len(dna) {
			return false
		}
	}
	return true
}
