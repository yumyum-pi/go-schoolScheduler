package generator

import (
	"fmt"
	"math/rand"
	"sort"
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

	(*p).Sort()
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

	shuffleNucleotide(&chromo)

	chromo.CheckEM2()
	chromo.HandleEM1()
	chromo.HandleEM2()
	chromo.CalFitness()

	chromo.GenCode = fmt.Sprintf("%02dN%02d:%04v", generation, index, chromo.nErr)
	return &chromo
}

func shuffleNucleotide(c *chromosome) {
	// loop though the nucleotides
	for nIndex := 0; nIndex < (*c).lSequence; nIndex += (*c).GeneSize {
		for n1 := 0; n1 < (*c).GeneSize; n1++ {
			n2 := rand.Intn((*c).GeneSize)
			(*c).SwapNucleotide(nIndex+n1, nIndex+n2)
		}
	}
}

// Print will write data to console
func (p *Population) Print() {
}

// PrintChromo will write data to console
func (p *Population) PrintChromo() {

}

// Next creates the next gene of chromosome
func (p *Population) Next(g int) {
	(*p).Wip()
	(*p).CrossOver()
	(*p).Mutate()
	(*p).Sort()
}

// CrossOver creates new chromosomes form the existing chromosomes
// by cross overing the genes
func (p *Population) CrossOver() {

}

// Mutate creates new chromosomes by changing
func (p *Population) Mutate() {

}

// Wip will delete
func (p *Population) Wip() {
	nc := make([]chromosome, pSize)
	copy(nc, (*p).P[:p2Size])
	copy((*p).P[:], nc)

}

// Sort will sort the data by fitness
func (p *Population) Sort() {
	nc := (*p).P
	sort.Slice(nc[:], func(p, q int) bool {
		return nc[p].Fitness > nc[q].Fitness
	})

	(*p).P = nc
}

// crossOver creates new nucleotide sequence by exchanging genes between two
// nucleotide sequence
func crossOver(s0, s1 *[]byte, gSize int) (*[]byte, *[]byte) {
	sl := len((*s0))
	// loop through gene
	s3 := make([]byte, 0, sl)
	s4 := make([]byte, 0, sl)

	flip := false
	for gIndex := 0; gIndex < sl; gIndex += gSize {
		geneEnd := gIndex + gSize

		if flip {
			s3 = append(s3, (*s0)[gIndex:geneEnd]...)
			s4 = append(s4, (*s1)[gIndex:geneEnd]...)
		} else {
			s3 = append(s3, (*s1)[gIndex:geneEnd]...)
			s4 = append(s4, (*s0)[gIndex:geneEnd]...)
		}
		flip = !flip
	}
	return &s3, &s4
}
