package main

import (
	"regexp"
)

func matchValidate(dna []string) bool {
	re := regexp.MustCompile("^(AAAA)|(GGGG)|(TTTT)|(CCCC)$")

	for i := 0; i < len(dna); i++ {
		if re.MatchString(dna[i]) {
			return true
			break
		}
	}
	return false
}
