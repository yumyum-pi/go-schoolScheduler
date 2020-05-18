package generator

import "fmt"

var empty = []byte{}

// Chromosome is the collection of TimeTable of the whole school.
// arrange in an slice of period.
type chromosome struct {
	GenCode     string // code of the generation
	GeneSize    int    // size of a gene sequence - class
	Sequence    []byte // sequence of nucleotide - period
	lSequence   int    // length of the sequence
	ErrSequence []byte // slice of conflicting nucleotides - periods
	nErr        int    // no of error
	Fitness     int    // fitness of the chromosome
	NDist       []byte // distribution of each nucleotide
}

// SwapNucleotide change the positions of nucleotide in the sequence
// and the error
func (c *chromosome) SwapNucleotide(sIndex0, sIndex1 int) {
	n0 := (*c).Sequence[sIndex0]
	n1 := (*c).Sequence[sIndex1]
	p0 := sIndex0 % (*c).GeneSize
	p1 := sIndex1 % (*c).GeneSize
	(*c).Sequence[sIndex0], (*c).Sequence[sIndex1] = (*c).Sequence[sIndex1], (*c).Sequence[sIndex0]
	(*c).ErrSequence[sIndex0], (*c).ErrSequence[sIndex1] = (*c).ErrSequence[sIndex1], (*c).ErrSequence[sIndex0]

	// swap the NDist
	dIndex0 := (int(n0-1) * (*c).GeneSize)
	dIndex1 := (int(n1-1) * (*c).GeneSize)
	// remove reduce from current position
	(*c).NDist[dIndex0+p0]--
	(*c).NDist[dIndex1+p1]--

	// add to current position
	(*c).NDist[dIndex0+p1]++
	(*c).NDist[dIndex1+p0]++
}

// illegalMutation checks for unwanted mutation cause by badly written code.
// The function take two nucleotide sequences(points of byte slice) - s1 && ns2
// and compared ns2 with s1 and throws error if :
//  - s1 and ns2 don't had the same nucleotide types
//  - quantities of nucleotide types not are equal
func illegalMutation(s0, s1 *[]byte, gSize int) error {
	// store length
	s0l := len(*s0)
	s1l := len((*s1))

	// check the lengths
	if s0l != s1l {
		return fmt.Errorf("s1 and ns2 don't have equal length")
	}

	// variables to store values to avoid reassigning
	// maps to store quantity of each nucleotides type present in a
	// genes of the respective sequence
	geneMap0 := make(map[byte]int, gSize)
	geneMap1 := make(map[byte]int, gSize)

	index := 0        // index of a nucleotide in the sequence
	var n0, n1 byte   // a single nucleotide
	var q0, q1 int    // quantity of a nucleotide type
	var ok0, ok1 bool // if nucleotide exists in the map

	// iterate over each gene to check all the nucleotides
	for i := 0; i < s0l; i += gSize {
		// assign all the nucleotides type and their quantity in the
		// gene to the respective maps
		for j := 0; j < gSize; j++ {
			index = i + j // calculate the index

			// nucleotide type in s0
			n0 = (*s0)[index]      // get the nucleotide at the index
			q0, ok0 = geneMap0[n0] // check if nucleotide exist in the map
			if !ok0 {
				q0 = 0 // assign initial value
			}
			q0++              // increase quantity of the nucleotide type
			geneMap0[n0] = q0 // reassign to the map

			// nucleotide type in s1
			n1 = (*s1)[index]      // get the nucleotide at the index
			q1, ok1 = geneMap1[n1] // check if nucleotide exist in the map
			if !ok1 {
				q1 = 0 // assign initial value
			}
			q1++              // increase quantity of the nucleotide type
			geneMap1[n1] = q1 // reassign to the map
		}

		// evalute the maps of the current gene
		for n0, q0 = range geneMap0 {
			// check if nucleotide type exist in geneMap2
			q1, ok1 = geneMap1[n0]
			if !ok1 {
				return fmt.Errorf(
					"n=%v is not present in the new chromosome",
					n0,
				)
			}
			// check if nucleotide type in geneMap2 has the same quantity
			if q0 != q1 {
				return fmt.Errorf(
					"n=%v quantity is not valid in the new chromosome."+
						"Was expecting q1=%v but got q2=%v at gene=%v",
					n0,
					q0,
					q1,
					i/gSize,
				)
			}
		}
	}
	return nil
}

// MatchN check nucleotides at the given position in each gene
// and returns the index of the matching nucleotide . If no match is
// found then -1 is returned
func (c *chromosome) MatchN(sIndex0 int) (sIndex1 int) {
	n := (*c).Sequence[sIndex0]

	// calculate gene position
	p := sIndex0 % (*c).GeneSize

	// loop though each gene
	for gIndex := 0; gIndex < (*c).lSequence; gIndex += (*c).GeneSize {
		sIndex1 = gIndex + p // calc the sequence index

		// check for matching nucleotide
		// avoid same sequence index
		if (*c).Sequence[sIndex1] == n && sIndex0 != sIndex1 {
			return
		}
	}
	return -1
}

// CheckEM1 (Check Error Method 1) checks for matching nucleotides in
// each gene position and updates the list of ErrIndexL
func (c *chromosome) CheckEM1() {
	(*c).nErr = 0
	// loop through each element
	for sIndex := 0; sIndex < (*c).lSequence; sIndex++ {
		// check conflit
		if n := (*c).MatchN(sIndex); n != -1 {
			//adding to error list
			(*c).ErrSequence[sIndex] = (*c).Sequence[sIndex]
			(*c).nErr++
		}

	}
}

// CheckEM2 (Check Error Method 1) checks for matching nucleotides in
// each gene position and updates the list of ErrIndexL
func (c *chromosome) CheckEM2() {
	(*c).nErr = 0
	// create a new empty byte slice to store error nucleotides
	(*c).ErrSequence = make([]byte, (*c).lSequence, (*c).lSequence)
	nGene := (*c).lSequence / (*c).GeneSize

	var sIndex0, sIndex1 int // store index for matching nucleotide
	var n0, n1 byte          // storing nucleotide of each index
	found := false           // match found

	// loop through each gene
	for gIndex0 := 0; gIndex0 < nGene; gIndex0++ {
		// loop through each nucleotide in gene
		for j := 0; j < (*c).GeneSize; j++ {
			// assigning index 0 value
			sIndex0 = gIndex0*(*c).GeneSize + j
			n0 = (*c).Sequence[sIndex0]

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
				n1 = (*c).Sequence[sIndex1]

				// matching nucleotides
				if n0 == n1 {
					// match is found
					found = true
					// reassign the nucleotide at index1 to 0
					// to skip in next iterations
					(*c).ErrSequence[sIndex1] = n1
					(*c).nErr++
				}

			}
			// check if match is found
			if found {
				// add nucleotide at sequence index0 to error slice
				(*c).ErrSequence[sIndex0] = n0
				(*c).nErr++
			}
		}
	}
	return
}

// CheckEM3 (Check Error Method 1) checks for matching nucleotides in
// each gene position and updates the list of ErrIndexL
func (c *chromosome) CheckEM3() {
	(*c).nErr = 0
	var p, dIndex int
	var f byte
	// loop through each element
	for sIndex, n := range (*c).Sequence {
		p = sIndex % (*c).GeneSize
		dIndex = int(n-1) * (*c).GeneSize

		f = (*c).NDist[dIndex+p]

		if f > 1 {
			//adding to error list
			(*c).ErrSequence[sIndex] = n
			(*c).nErr++
		}
	}
}

// HandleEM1(Handle Error Method 1) tried to correct the overlapping
// nucleotides by interchaning the position of the error nucleotides which
// are in the same genes and don't overlap anymore
func (c *chromosome) HandleEM1() {
	var n0, n1 byte

	// store values
	var g0, g1 int // gene index
	//var b0, b1 int // swap safe

	var sIndex0, sIndex1 int
	// loop through each error index
	for sIndex0, n0 = range (*c).ErrSequence {
		g0 = sIndex0 / (*c).GeneSize
		g1 = (g0 + 1) * (*c).GeneSize

		// skip if nucleotide is error free
		if n0 == 0 {
			continue
		}

		// loop through all the index in the current gene
		// begining from the current gene index start
		for sIndex1 = sIndex0; sIndex1 < g1; sIndex1++ {
			n1 = (*c).ErrSequence[sIndex1]

			// skip if
			// nucleotide is error free
			// index match
			if n1 == 0 || sIndex1 == sIndex0 {
				continue
			}

			// check if swap is use full
			if (*c).CheckSafeSwap(sIndex0, sIndex1) {
				break
			}
		}
	}
}

// HandleEM2(Handle Error Method 2) tried to correct the overlapping
// nucleotides by interchaning the position of the nucleotides which
// are in the same genes and don't overlap anymore
func (c *chromosome) HandleEM2() error {
	// store variable
	var g, sIndex1 int // gene index

	// loop through error index list
	for sIndex0, n := range (*c).ErrSequence {
		g = sIndex0 / (*c).GeneSize
		if n == 0 {
			continue
		}
		for p := 0; p < (*c).GeneSize; p++ {
			sIndex1 = g*(*c).GeneSize + p

			// check if swap is use full
			if (*c).CheckSafeSwap(sIndex0, sIndex1) {
				break
			}
		}
	}
	return nil
}

// HandleEM3(Handle Error Method 2) tried to correct the overlapping
// nucleotides by interchaning the position of the nucleotides which
// are in the same genes and don't overlap anymore
func (c *chromosome) HandleEM3() {
	var n0, n1 byte          // nucleotide
	var dIndex0, dIndex1 int // index of nucleotide distribution
	var f0, f1 byte          // frequency of nucleotide
	var p0, p1 int           // position of nucleotides

	for gIndex := 0; gIndex < (*c).lSequence; gIndex += (*c).GeneSize {
		for p0 = 0; p0 < (*c).GeneSize; p0++ {
			n0 = (*c).Sequence[gIndex+p0]
			dIndex0 = int(n0-1) * (*c).GeneSize

			f0 = (*c).NDist[dIndex0+p0] // get the frequency at p0

			// skip if nucleotide has not overlaps
			if f0 < 2 {
				continue
			}
			for p1 = 0; p1 < (*c).GeneSize; p1++ {
				if p1 == p0 {
					continue
				}
				n1 = (*c).Sequence[gIndex+p1]
				dIndex1 = int(n1-1) * (*c).GeneSize

				// check if nucleotides have overlaps
				f0 = (*c).NDist[dIndex0+p1] // n0 at p1
				f1 = (*c).NDist[dIndex1+p0] // n1 at p2

				// swap positions if nucleotide has not overlaps
				if f1 == 0 && f0 == 0 {

					(*c).Sequence[gIndex+p0], (*c).Sequence[gIndex+p1] = (*c).Sequence[gIndex+p1], (*c).Sequence[gIndex+p0]
					(*c).NDist[dIndex0+p0]--
					(*c).NDist[dIndex1+p1]--

					(*c).NDist[dIndex0+p1]++
					(*c).NDist[dIndex1+p0]++

					//(*c).SwapNucleotide((gIndex + p0), (gIndex + p1))
					break
				}
			}
		}
	}
}

// CheckSafeSwap takes two variable in the same gene and checks if swaping
// their postion will resolve problem and return int. Meaning of returned int
//  (-1) - can not resolve conflict
//	(0) - can resolve conflict
//  (0 <) - can resolve conflict of given and return index
func (c *chromosome) CheckSafeSwap(sIndex0, sIndex1 int) bool {
	var b0, b1 bool = true, true

	// get nucleotide
	n0 := (*c).Sequence[sIndex0] // sequence index 0
	n1 := (*c).Sequence[sIndex1] // sequence index 1

	// position of nucleotide
	p0 := sIndex0 % (*c).GeneSize // sequence index 0
	p1 := sIndex1 % (*c).GeneSize // sequence index 1

	gIndex0 := sIndex0 / (*c).GeneSize // current gene index

	var tn0, tn1 byte // nucleotide at test index

	// loop through each gene
	for gIndex := 0; gIndex < (*c).lSequence; gIndex += (*c).GeneSize {
		// skip if the gene index is same
		if gIndex == gIndex0 {
			continue
		}

		// get nucleotides with the same position in different gene
		tn0 = (*c).Sequence[gIndex+p1]
		tn1 = (*c).Sequence[gIndex+p0]

		// if b1 has not issues
		// check conflict at p1 of n0
		if b0 {
			// check if n0 has a match at p1 position
			if n0 == tn0 {
				b0 = false
			}
		}

		// if b1 has not issues
		// check conflict at p0 of n1
		if b1 {
			if n1 == tn1 {
				b1 = false
			}
		}
	}
	// no conflicting match is found
	if b0 && b1 {
		(*c).SwapNucleotide(sIndex0, sIndex1)
		// remove the error from the list
		(*c).ErrSequence[sIndex0] = 0
		(*c).ErrSequence[sIndex1] = 0
		return true
	}
	return false
}

// Print writes out to stout
func (c *chromosome) Print(detail bool) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorRed := "\033[1;31m"

	index := 0 // index of a nucleotide in the sequence

	fmt.Printf(
		"genCode=%v\tgeneSize=%v\tnErr=%v\tFitness=%v\n",
		(*c).GenCode,
		(*c).GeneSize,
		(*c).nErr,
		(*c).Fitness,
	)
	if detail {
		for i := 0; i < (*c).lSequence; i += (*c).GeneSize {
			fmt.Printf("%3v[ ", i/(*c).GeneSize)
			for j := 0; j < (*c).GeneSize; j++ {
				index = i + j
				if j%8 == 0 && j != 0 {
					fmt.Printf("%v| ", string(colorReset))
				}
				// check if error
				if (*c).ErrSequence[index] != 0 {
					fmt.Printf("%v%02v%v ", string(colorRed), (*c).Sequence[index], string(colorReset))
					continue

				}
				fmt.Printf("%v%02v ", string(colorGreen), (*c).Sequence[index])
			}
			fmt.Printf("%v]\n", string(colorReset))
		}
	}

}

func (c *chromosome) CalFitness() {
	// update the check list
	(*c).CheckEM2()

	// calculate the fitness by error
	if (*c).nErr != 0 {
		// negative fitness for error
		(*c).Fitness = (*c).lSequence - (5 * (*c).nErr)
	} else {
		// additional fitness for 0 error
		(*c).Fitness = (*c).lSequence * 2
	}
	// make map of nucleotide that contains the slice of bool which
	// represents the gene position they are in.
	mapN := make(map[byte][]bool)
	// store values
	var ok bool  // is present in the map
	var p int    // position in the gene
	var q []bool // list of position in the gene

	// loop though each nucleotide in the sequence
	for sIndex, n := range (*c).Sequence {
		// calculate the position
		p = sIndex % (*c).GeneSize

		q, ok = mapN[n]
		if !ok {
			q = make([]bool, (*c).GeneSize)
		}
		// assgin the position in slice
		q[p] = true
		// reassign the bool slice to the map
		mapN[n] = q

	}

	i := 0 // sum of all the variation
	for _, q := range mapN {
		sumDiff := 0         // sum of all the difference
		var sumDiff2 float32 // sum of diff^2
		dif := 0             // difference between 2 position in the gene
		var d float32        // total no. of diff
		l := 0               // last position
		for j, b := range q {
			if b {
				dif = j - l // subtract the current position with last
				l = j       // make current position current position
				sumDiff += dif
				sumDiff2 += float32(dif * dif)
				d++ // increse the diff no.
			}
		}

		m := float32(sumDiff) / d     // calculate mean
		v := (sumDiff2 / d) - (m * m) // calculate variance
		i += int(v * 100)             // decrease decimal point
	}

	(*c).Fitness -= i
}

// PrintSequence writes to the stout
func PrintSequence(s0 *[]byte, gSize int) {
	l := len(*s0)

	index := 0 // index of a nucleotide in the sequence

	for i := 0; i < l; i += gSize {
		fmt.Printf("%2v[ ", i/gSize)
		for j := 0; j < gSize; j++ {
			index = i + j // calculate the index
			if j%8 == 0 && j != 0 {
				fmt.Printf("| ")
			}
			fmt.Printf("%02v ", (*s0)[index])
		}
		fmt.Printf("]\n")
	}
}
