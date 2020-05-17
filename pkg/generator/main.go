package generator

import "fmt"

const nGeneration = 16

// Start begin the generating process
func Start(s0 *[]byte, geneSize int) (*[]byte, error) {
	var p Population
	p.Init(s0, geneSize)
	for g := 0; g < nGeneration; g++ {
		p.Next(g + 1)
	}
	for _, c := range p.P {
		if c.nErr == 0 {
			return &c.Sequence, nil
		}
	}

	return &p.P[0].Sequence, fmt.Errorf("Error=%04v ng=%03v", p.P[0].nErr, nGeneration)
}
