package database

type status struct {
	CountMutant float32 `json:"count_mutant_dna"`
	CountHuman  float32 `json:"count_human_dna"`
	Ratio       float32 `json:"ratio"`
}
