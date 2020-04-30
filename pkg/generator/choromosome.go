package generator

import "fmt"

var empty = []byte{}

// Chromosome is the collection of TimeTable of the whole school.
// arrange in an slice of period.
type chromosome struct {
	GeneSize    int    // size of a gene sequence - class
	Nucleotides []byte // sequence of nucleotide - period
	Name        string // name of genes
	ErrIDs      []int  // slice of conflicting nucleotides - periods
	Fitness     int    // fitness of the chromosome

}

// Length return the length of genes
func (c *chromosome) Length() int {
	return len((*c).Nucleotides)
}

// SwapNucleotide change the positions of nucleotide in the dna sequence
func (c *chromosome) SwapNucleotide(n1, n2 int) {
	(*c).Nucleotides[n1], (*c).Nucleotides[n2] = (*c).Nucleotides[n2], (*c).Nucleotides[n1]
}

func illegalMutation(ns1, ns2 *[]byte, gSize int) error {
	ns1l := len(*ns1)   // length of (*ns1)
	ns2l := len((*ns2)) // length of newChromosome

	// check the length of nucleotides
	if ns1l != ns2l {
		return fmt.Errorf("ns1l and (*ns2) don't have equal length")
	}
	ns1GeneMap := make(map[byte]int, gSize)
	ns2GeneMap := make(map[byte]int, gSize)
	index := 0  // index of the nucleotide sequence
	var n byte  // a single nucleotide
	q := 0      // quantity of a type of nucleotide
	ok := false // if nucleotide is presen in the map
	for i := 0; i < ns1l; i += gSize {
		// check each gene
		// each gene should have the same quantity of nucleotides of different types
		// when when compared to (*ns1)'s gene
		for j := 0; j < gSize; j++ {
			// calculate the quantity of each
			// type of nucleotide in ns1
			index = i + j
			n = (*ns1)[index]
			q, ok = ns1GeneMap[n]
			if !ok {
				q = 0
			}
			q++
			ns1GeneMap[n] = q

			// calculate the quantity of each
			// type of nucleotide in new chromosome
			n = (*ns2)[index]
			q, ok = ns2GeneMap[n]
			if !ok {
				q = 0
			}
			q++
			ns2GeneMap[n] = q
		}

		for n, q1 := range ns1GeneMap {
			q, ok = ns2GeneMap[n]
			if !ok {
				return fmt.Errorf("n=%v is not present in the new chromosome", n)
			}

			if q != q1 {
				return fmt.Errorf("n=%v quantity is not valid in the new chromosome", n)
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
