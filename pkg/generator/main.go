package generator

import "fmt"

const nGeneration = 16

// Start begin the generating process
func Start(s0 *[]byte, geneSize int) (*[]byte, int, error) {
	var p Population
	p.Init(s0, geneSize)
	for g := 0; g < nGeneration; g++ {
		p.Next(g + 1)
	}
	for _, c := range p.P {
		if c.nErr == 0 {
			return &c.Sequence, 0, nil
		}
	}

	return &p.P[0].Sequence, p.P[0].nErr, fmt.Errorf("Error=%04v ng=%03v", p.P[0].nErr, nGeneration)
}

// create nucleotide
func nDistrubution(s0 *[]byte, gSize int, NNType int) (nDist *[]byte) {
	var n int
	dist := make([]byte, NNType*gSize)
	// loop though each nucleotide in the sequence
	for sIndex := 0; sIndex < len(*s0); sIndex += gSize {
		// loop through each gene
		for p := 0; p < gSize; p++ {
			n = int((*s0)[sIndex+p]) - 1
			dist[(n*gSize)+p]++
		}
	}

	return &dist
}

func printNDistribution(dist *[]byte, gSize int) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorRed := "\033[1;31m"

	// different color option
	var n byte
	var t int

	// loop through each teacher
	for i := 0; i < len(*dist); i += gSize {
		t = i / gSize
		t++
		fmt.Printf("%03v[", t)
		for p := 0; p < gSize; p++ {
			k := p % 8
			if k == 0 && p != 0 {
				fmt.Printf(" |")
			}
			n = (*dist)[i+p]
			switch n {
			case 0:
				fmt.Printf(" --")
			case 1:
				fmt.Printf(" %v01%v", string(colorGreen), string(colorReset))
			default:
				fmt.Printf(" %v%02v%v", string(colorRed), n, string(colorReset))
			}
		}
		fmt.Printf(" ]\n")
	}

}
