package generator

import "fmt"

var empty = []byte{}

// Chromosome is the collection of TimeTable of the whole school.
// arrange in an slice of period.
type chromosome struct {
	GenCode   string // code of the generation
	GeneSize  int    // size of a gene sequence - class
	Sequence  []byte // sequence of nucleotide - period
	ErrIndexL []int  // slice of conflicting nucleotides index - periods
	Fitness   int    // fitness of the chromosome

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

// MatchN check nucleotides at the given position in each gene
// and returns the index of the matching nucleotide . If no match is
// found then -1 is returned
func (c *chromosome) MatchN(sIndex int) (sIndex2 int) {
	l := (*c).Length()
	n := (*c).Sequence[sIndex]
	// calculate gene index
	gIndex := sIndex % (*c).GeneSize

	for i := 0; i < l; i += (*c).GeneSize {
		sIndex2 = i + gIndex
		if (*c).Sequence[sIndex2] == n && sIndex != sIndex2 {
			return
		}
	}
	return -1
}

// CheckEM1 (Check Error Method 1) checks for matching nucleotides in each
// gene position and updates the list of ErrIndexL
func (c *chromosome) CheckEM1() {
	var err []int // list of error index
	l := (*c).Length()

	// loop through each element
	for sIndex := 0; sIndex < l; sIndex++ {
		// check conflit
		if n := (*c).MatchN(sIndex); n != -1 {
			// the sIndex to the error list
			err = append(err, sIndex)
		}

	}
	// update the chromosome error list
	(*c).ErrIndexL = err
	return
}

// CheckEM2 (Check Error Method 1) checks for matching nucleotides in each
// gene position and updates the list of ErrIndexL
func (c *chromosome) CheckEM2() {
	var err []int // list of error index
	l := (*c).Length()
	nGene := l / (*c).GeneSize

	// new sequence to edit
	s := append([]byte{}, (*c).Sequence...)

	sIndex0, sIndex1 := 0, 0   // store index for matching nucleotide
	n0, n1 := byte(0), byte(0) // storing nucleotide of each index
	found := false             // match found

	// loop through each gene
	for gIndex0 := 0; gIndex0 < nGene; gIndex0++ {
		// loop through each nucleotide in gene
		for j := 0; j < (*c).GeneSize; j++ {
			// assigning index 0 value
			sIndex0 = gIndex0*(*c).GeneSize + j
			n0 = s[sIndex0]

			// skip if last gene or if n0 is 0
			if gIndex0 == nGene-1 || n0 == 0 {
				continue
			}

			found = false

			// loop through the next generations and find the match
			// in the same pID
			for gIndex1 := gIndex0 + 1; gIndex1 < nGene; gIndex1++ {
				// assigning index 1
				sIndex1 = gIndex1*(*c).GeneSize + j
				n1 = s[sIndex1]

				// matching nucleotides
				if n0 == n1 {
					// match is found
					found = true
					// reassign the nucleotide at index1 to 0
					// to skip in next iterations
					s[sIndex1] = 0

					// add the index to the error list
					err = append(err, sIndex1)
				}

			}
			// check if match is found
			if found {
				// reassign the nucleotide at index1 to 0
				// to skip in next iterations
				s[sIndex0] = 0

				// add the index to error list
				err = append(err, sIndex0)
			}
		}
	}
	// update the chromosome error list
	(*c).ErrIndexL = err
	return
}

// Print writes out to stout
func (c *chromosome) Print() {
	fmt.Printf(
		"genCode=%v\tgeneSize=%v\tnErr=%v\tFitness=%v\n",
		(*c).GenCode,
		(*c).GeneSize,
		len((*c).ErrIndexL),
		(*c).Fitness,
	)

	PrintSequence(&(*c).Sequence, (*c).GeneSize)
}

func (c *chromosome) PrintError() {
	l := (*c).Length()
	index := 0 // index of a nucleotide in the sequence

	nextIndex := 0
	el := len((*c).ErrIndexL)
	for i := 0; i < l; i += (*c).GeneSize {
		fmt.Printf("%2v[ ", i/(*c).GeneSize)
		for j := 0; j < (*c).GeneSize; j++ {
			index = i + j
			if nextIndex >= el {
				fmt.Printf("-- ")
				continue
			}
			n := (*c).ErrIndexL[nextIndex]
			if index == n {
				nextIndex++
				fmt.Printf("%2v ", (*c).Sequence[n])
				continue
			}
			fmt.Printf("-- ")
		}
		fmt.Printf("]\n")

	}
}

func deleteEml(s *[]byte, i int) {
	l := len(*s) - 1
	(*s)[i] = (*s)[l] // Copy last element to index i.
	(*s) = (*s)[:l]
}

// PrintSequence writes to the stout
func PrintSequence(s0 *[]byte, gSize int) {
	l := len(*s0)

	index := 0 // index of a nucleotide in the sequence

	for i := 0; i < l; i += gSize {
		fmt.Printf("%3v[ ", i/gSize)
		for j := 0; j < gSize; j++ {
			index = i + j // calculate the index
			fmt.Printf("%3v ", (*s0)[index])
		}
		fmt.Printf("]\n")
	}
}
