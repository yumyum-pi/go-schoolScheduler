package models

import "fmt"

const byteSize = 32

// Decode nucleotide sequence data from package
func Decode(pkgs *[][]byte, gSize int32) (*[]byte, int, error) {
	// calculate the capacity
	l := len((*pkgs))
	// get the length of the last byte
	lb := len((*pkgs)[l-1])
	cap := (l - 1) * 32
	cap += lb

	s0 := make([]byte, 0, cap)
	// loop through each package byte
	for _, pkg := range *pkgs {
		s0 = append(s0, pkg...)
	}

	if e := checkData(&s0, int(gSize)); e != nil {
		return nil, 0, e
	}

	return &s0, int(gSize), nil
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

// Encode nucleotide sequence data to packages
func Encode(pp *[]byte) (pkgs *[][]byte) {
	l := len(*pp)
	nPPkg := (byteSize) // no of period per package
	tnPPkg := l / nPPkg // total no. of package
	rem := l % nPPkg    // check if their is a remainder

	// create package array
	pkg := make([][]byte, tnPPkg)

	//p := make([]byte, nPPkg) // package
	var pi int = 0 // period index
	for i := 0; i < tnPPkg; i++ {
		// create package array
		p := make([]byte, nPPkg) // package
		for j := 0; j < nPPkg; j++ {
			pi = (i * nPPkg) + j
			//fmt.Println(pi)
			p[j] = (*pp)[pi] // get the byes of periods
			// add the period byte to the package
		}
		pkg[i] = p
	}
	// add remaining
	if rem != 0 {
		//	tnPPkg
		p := make([]byte, rem)
		for j := 0; j < rem; j++ {
			pi = (tnPPkg * byteSize) + j
			p[j] = (*pp)[pi] // get the byes of periods
			// add the period byte to the package
		}
		pkg = append(pkg, p)
	}

	return &pkg
}
