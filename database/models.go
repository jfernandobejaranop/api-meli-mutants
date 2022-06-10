package database

type status struct {
	CountMutant int `json:"count_mutant_dna"`
	CountHuman  int `json:"count_human_dna"`
	Ratio       int `json:"ratio"`
}
