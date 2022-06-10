package main

import (
	"fmt"
	"strings"
)

var arrC [][]string
var letra string

func HorizontalArray(dna []string) bool {
	fmt.Println("Horizontal : ", dna)
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
	fmt.Println("Vertical : ", arrB)
	return matchValidate(arrB)
}

func RightDiagonalArray(dna []string) bool {

	var arrG [12][12]string
	var arrD []string
	var arrH []string
	lenDna := len(dna)

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
		arrD = append(arrD, letra)
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
		arrD[fil] = arrD[fil] + letra
		letra = ""
	}

	//Pivoteo matriz
	for fil := 0; fil < len(arrD); fil++ {
		splits := strings.Split(arrD[fil], "")
		for col := range splits {
			if !(splits[col] == "") {
				arrG[col+fil][fil] = splits[col]
			}
		}
	}

	// Pivote final
	for fil := range arrG {
		for col := 0; col < len(arrG[fil]); col++ {
			letra = letra + arrG[fil][col]
		}
		arrH = append(arrH, letra)
		letra = ""
	}

	fmt.Println("Diagonal derecha : ", arrH)
	return matchValidate(arrH)
}

func LeftDiagonalArray(dna []string) bool {

	var arrI [12][12]string
	var arrJ []string
	var arrK []string
	var cont int = 0
	lenDna := len(dna)

	//PRIMERA PARTE MATRIZ
	for fil := range arrC {
		for col := 0; col < len(arrC[fil]); col++ {
			if col >= len(arrC[fil])-fil {
				break
			}
			letra = letra + arrC[col][fil]
		}
		arrJ = append(arrJ, letra)
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
		arrJ[cont] = arrJ[cont] + letra
		letra = ""
		cont++
	}

	//PIVOTEO MATRIZ COMPLETA
	for fil := 0; fil < len(arrJ); fil++ {
		splits := strings.Split(arrJ[fil], "")
		for col := range splits {
			if !(splits[col] == "") {
				arrI[col+fil][fil] = splits[col]
			}
		}
	}

	// PIVOTE FINAL
	for fil := range arrI {
		for col := 0; col < len(arrI[fil]); col++ {
			letra = letra + arrI[fil][col]
		}
		arrK = append(arrK, letra)
		letra = ""
	}

	fmt.Println("Diagonal izquierda : ", arrK)
	return matchValidate(arrK)
}
