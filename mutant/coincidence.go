package mutant

import (
	"regexp"
)

func MatchValidateDNA(dna []string) bool {
	re := regexp.MustCompile("(AAAA)|(GGGG)|(TTTT)|(CCCC)")
	for i := 0; i < len(dna); i++ {
		if re.MatchString(dna[i]) {
			return true
		}
	}
	return false
}
