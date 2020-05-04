package generator

import (
	"fmt"
	"math/rand"
)

const pSize = 32         // population size
const p2Size = pSize / 2 // half of population size
const p4Size = pSize / 4 // quater of population size

// Population is a fixed array of size 64 containing chromosomes
type Population struct {
	P            [pSize]chromosome
	totalError   int
	minErr       int
	maxErr       int
	errChrmsmIDs []int
	ns0          *[]byte
}

// Init create the initail population
func (p *Population) Init(ns0 *[]byte, gSize int) {
	//nL := len(*ns0) // length of nucleotides
	(*p).ns0 = ns0
	// make a new chromosome var

	for i := 0; i < pSize; i++ {
		(*p).P[i] = *newChromo(ns0, gSize, 0, i)
	}
}

// newChromo creates a new chromosome with the given sequence of nucleotides
func newChromo(ns0 *[]byte, gSize, generation, index int) *chromosome {
	// make a new chromosome
	var chromo chromosome
	chromo.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	chromo.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	chromo.ErrSequence = make([]byte, nL, nL)
	chromo.lSequence = nL

	// loop though the nucleotides
	for nIndex := 0; nIndex < nL; nIndex += gSize {
		for n1 := 0; n1 < gSize; n1++ {
			n2 := rand.Intn(gSize)
			chromo.SwapNucleotide(nIndex+n1, nIndex+n2)
		}
	}

	chromo.CheckEM2()
	chromo.HandleEM1()
	chromo.HandleEM2()
	chromo.CalFitness()

	chromo.GenCode = fmt.Sprintf("%02dN%02d:%04v", generation, index, chromo.nErr)
	return &chromo
}

// Print will write data to console
func (p *Population) Print() {
}

// PrintChromo will write data to console
func (p *Population) PrintChromo() {

}

// Next creates the next gene of chromosome
func (p *Population) Next() {
	(*p).Wip()
	(*p).CrossOver()
	(*p).Sort()
}

// CrossOver creates new chromosomes form the existing chromosomes
// by cross overing the genes
func (p *Population) CrossOver() {

}

// Wip will delete
func (p *Population) Wip() {
	nc := make([]chromosome, pSize)
	copy(nc, (*p).P[:p2Size])
	copy((*p).P[:], nc)

}

// Sort will sort the data by fitness
func (p *Population) Sort() {

}

// crossOver creates new nucleotide sequence by exchanging genes between two
// nucleotide sequence
func crossOver() {

}
