package main

import (
	"fmt"
	"regexp"
)

func matchValidate(dna []string) bool {
	re := regexp.MustCompile("^(AAAA)|(GGGG)|(TTTT)|(CCCC)$")

	for i := 0; i < len(dna); i++ {
		if re.MatchString(dna[i]) {
			fmt.Println("entro : ", dna[i])
			return true
		}
	}
	return false
}
