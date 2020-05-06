package models

import "fmt"

// Decode nucleotide sequence data
func (s *SequencePkgs) Decode() (*[]byte, int, error) {
	// calculate the capacity
	l := len((*s).Pkgs)
	// get the length of the last byte
	lb := len((*s).Pkgs[l-1])
	cap := (l - 1) * 32
	cap += lb

	s0 := make([]byte, 0, cap)
	// loop through each package byte
	for _, pkg := range (*s).Pkgs {
		s0 = append(s0, pkg...)
	}

	if e := checkData(&s0, int((*s).GSize)); e != nil {
		return nil, 0, e
	}

	return &s0, int((*s).GSize), nil
}

// checkData checks if the quantity of each nucleotide type is < geneSize
// else thro error
func checkData(s0 *[]byte, geneSize int) error {
	s0l := len(*s0)

	// total period counter of a teacher
	geneMap := make(map[byte]int)

	// variables to store values to avoid reassigning
	var n byte  // single nucleotide
	q := 0      // quantity of a nucleotide type
	ok := false // if nucleotide exists in the map

	// iterate over each gene to check all the nucleotides
	for i := 0; i < s0l; i += geneSize {
		// assign all the nucleotides type and their quantity in
		// the gene to the respective maps
		for j := 0; j < geneSize; j++ {
			n = (*s0)[(i + j)] // get the nucleotide at the index

			q, ok = geneMap[n] // check if nucleotide exist in the map
			if !ok {
				q = 0 // assign initial value
			}
			q++            // increase quantity of the nucleotide type
			geneMap[n] = q // reassign to the map
		}

	}
	// evalute the maps
	for _, q := range geneMap {
		// check the quantify of the nucleotide
		if q >= geneSize {
			return fmt.Errorf("> Error: Data received is invalid")
		}
	}

	return nil
}
