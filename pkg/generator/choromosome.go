package generator

import "fmt"

var empty = []byte{}

// Chromosome is the collection of TimeTable of the whole school.
// arrange in an slice of period.
type chromosome struct {
	GeneSize int    // size of a gene sequence - class
	Sequence []byte // sequence of nucleotide - period
	Name     string // name of genes
	ErrIDs   []int  // slice of conflicting nucleotides - periods
	Fitness  int    // fitness of the chromosome

}

// Length return the length of sequence
func (c *chromosome) Length() int {
	return len((*c).Sequence)
}

// SwapNucleotide change the positions of nucleotide in the sequence
func (c *chromosome) SwapNucleotide(n1, n2 int) {
	(*c).Sequence[n1], (*c).Sequence[n2] = (*c).Sequence[n2], (*c).Sequence[n1]
}

// illegalMutation checks for unwanted mutation cause by badly written code.
// The function take two nucleotide sequences(points of byte slice) - ns1 && ns2
// and compared ns2 with ns1 and throws error if :
//  - ns1 and ns2 had the same types of nucleotides
//  - quantities of nucleotides are equal
func illegalMutation(ns1, ns2 *[]byte, gSize int) error {
	// store length
	ns1l := len(*ns1)
	ns2l := len((*ns2))

	// check the lengths
	if ns1l != ns2l {
		return fmt.Errorf("ns1 and ns2 don't have equal length")
	}

	// variables to store values to avoid reassigning
	// maps to store quantity of each nucleotides type present in a genes
	// of the respective sequence
	geneMap1 := make(map[byte]int, gSize)
	geneMap2 := make(map[byte]int, gSize)

	index := 0  // index of a nucleotide in the sequence
	var n byte  // a single nucleotide
	q := 0      // quantity of a nucleotide type
	ok := false // if nucleotide exists in the map

	// iterate over each gene to check all the nucleotides
	for i := 0; i < ns1l; i += gSize {
		// assign all the nucleotides type and their quantity in the
		// gene to the respective maps
		for j := 0; j < gSize; j++ {
			index = i + j // calculate the index

			// nucleotide type in ns1
			n = (*ns1)[index]   // get the nucleotide at the index
			q, ok = geneMap1[n] // check if nucleotide exist in the map
			if !ok {
				q = 0 // assign initial value
			}
			q++             // increase quantity of the nucleotide type
			geneMap1[n] = q // reassign to the map

			// nucleotide type in ns2
			n = (*ns2)[index]   // get the nucleotide at the index
			q, ok = geneMap2[n] // check if nucleotide exist in the map
			if !ok {
				q = 0 // assign initial value
			}
			q++             // increase quantity of the nucleotide type
			geneMap2[n] = q // reassign to the map
		}

		// evalute the maps of the current gene
		for n, q1 := range geneMap1 {
			// check if nucleotide type exist in geneMap2
			q, ok = geneMap2[n]
			if !ok {
				return fmt.Errorf(
					"n=%v is not present in the new chromosome",
					n,
				)
			}
			// check if nucleotide type in geneMap2 has the same quantity
			if q != q1 {
				return fmt.Errorf(
					"n=%v quantity is not valid in the new chromosome",
					n,
				)
			}
		}
	}
	return nil
}

func deleteEml(s *[]byte, i int) {
	l := len(*s) - 1
	(*s)[i] = (*s)[l] // Copy last element to index i.
	(*s) = (*s)[:l]
}
