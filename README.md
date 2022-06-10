
# Technical challenge for MELI

Api rest that identifies mutants by the nitrogenous base of ADN 
_______________________________________________

### Consume
End Points:

## 1. Mutant service: 

End point whit detects if a human is mutant sending the DNA sequence through an HTTP POST with a Json which has the following format:

- **URL:** https://api-meli-mutants.herokuapp.com/mutant

- **Method:** POST

- **Body Params:**

  ```json
  {"dna": [[string], [string], [string], ...]}
  ```
 **Example:**

  ```json
  {"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}

  {"dna":["ATGT","CAGT","TTAT","AGAA"]}

  {"dna":["ATGTAAAC","CAGTTTGA","TTATGACC","AGAACCTA","ATGTCAGT","TTATAGAA","CAGTTTGA","TTATGACC"]}
  ```
  
- **Response:**
**Is Mutant** = http.status = 200 - OK
**Is Not Mutant** http.status = 403 - Forbidden
 
   
### 2. Stats:

End Point which returns statistics of the ADN checks

- **URL:** https://api-meli-mutants.herokuapp.com/stats

- **Method:** GET

- **Response:** 

  ```Json
  {
    "count_mutant_dna": [number],
    "count_human_dna": [number],
    "ratio": [number]
  }
  ```
_________________________________________________