package generator

import "fmt"

const nGeneration = 16

// Start begin the generating process
func Start(s0 *[]byte, geneSize, nNType int) (*[]byte, int, error) {
	nGene := len(*s0) / geneSize
	nGene /= 12
	nGene = nGene * nGene * nGene / 2
	nDist := nDistribution(s0, geneSize, nNType)
	p := CreatePopulation(s0, nDist, geneSize, nNType)
	p.Init()
	for g := 0; g < nGene; g++ {
		p.Next(g + 1)
	}
	for _, c := range p.P {
		if c.nErr == 0 {
			if er := illegalMutation(s0, &c.Sequence, geneSize); er != nil {
				fmt.Println(er)
			}
			return &c.Sequence, 0, nil
		}
	}
	if er := illegalMutation(s0, &p.P[0].Sequence, geneSize); er != nil {
		fmt.Println(er)
	}
	return &p.P[0].Sequence, p.P[0].nErr, fmt.Errorf("Error=%04v ng=%03v", p.P[0].nErr, nGene)
}

// create nucleotide
func nDistribution(s0 *[]byte, gSize, NNType int) (nDist *[]byte) {
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
