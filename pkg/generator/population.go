package generator

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

const pSize = 32                // population size
const p2Size = pSize / 2        // half of population size
const p4Size = pSize / 4        // quarter of population size
const p34Size = p4Size + p2Size // 3/4 of population size

// Population is a fixed array of size 64 containing chromosomes
type Population struct {
	P      []chromosome
	ns0    *[]byte
	nDist  *[]byte
	gSize  int
	nNType int
}

<<<<<<< HEAD
// Init create the initial population
func (p *Population) Init(ns0 *[]byte, gSize int) {
	// assign the initial values
	(*p).ns0 = ns0
	(*p).gSize = gSize

=======
// CreatePopulation return a new population
func CreatePopulation(ns0, nDist *[]byte, gSize, nNType int) *Population {
	chromosomes := make([]chromosome, pSize)
	return &Population{
		chromosomes,
		ns0,
		nDist,
		gSize,
		nNType,
	}
}

// Init create the initial population
func (p *Population) Init() {
>>>>>>> iss7
	var wg sync.WaitGroup

	for i := 0; i < pSize; i++ {
		wg.Add(1)

		// run the function concurrently
		go func(i int) {
<<<<<<< HEAD
			(*p).P[i] = *newChromo((*p).ns0, (*p).gSize, 0, i-p34Size)
=======
			(*p).P[i] = *newChromo((*p).ns0, (*p).nDist, (*p).gSize, 0, i-p34Size)
>>>>>>> iss7
			wg.Done()
		}(i)
	}

	// waiting for all the task to get over
	wg.Wait()

	// sort the population on fitness
	(*p).Sort()
}

// newChromo creates a new chromosome with the given sequence of nucleotides
func newChromo(ns0, nDist *[]byte, gSize, generation, index int) *chromosome {
	// make a new chromosome
	var chromo chromosome
	chromo.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	chromo.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	chromo.NDist = append((*nDist)[:0:0], (*nDist)...)
	chromo.ErrSequence = make([]byte, nL, nL)
	chromo.lSequence = nL

	shuffleNucleotide(&chromo)

	chromo.HandleEM3()
	chromo.HandleEM3()
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
func (p *Population) Print(details bool) {
	for i := 0; i < pSize; i++ {
		(*p).P[i].Print(details)
	}
}

// Next creates the next gene of chromosome
func (p *Population) Next(g int) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		(*p).CrossOver(g)
		wg.Done()
	}()

	go func() {
		(*p).Mutate(g)
		wg.Done()
	}()

	go func() {
		(*p).New(g)
		wg.Done()
	}()
	wg.Wait()

	(*p).Sort()
}

<<<<<<< HEAD
// Crossover creates new chromosomes form the existing chromosomes
=======
// CrossOver creates new chromosomes form the existing chromosomes
>>>>>>> iss7
// by mixing the genes of two chromosome
func (p *Population) CrossOver(g int) {
	var wg sync.WaitGroup

	// loop through 1/4 for the list in pair of two
	for i := 0; i < p4Size; i += 2 {
		wg.Add(1)
		go func(i int) {
			// get the new cross over sequences
			ns0, ns1 := crossOver(&(*p).P[i].Sequence, &(*p).P[i+1].Sequence, (*p).P[i].GeneSize)
<<<<<<< HEAD

=======
			nDist0 := nDistribution(ns0, p.gSize, p.nNType)
			nDist1 := nDistribution(ns1, p.gSize, p.nNType)
>>>>>>> iss7
			// create chromosomes
			var c0, c1 chromosome
			c0.GeneSize = (*p).gSize
			c1.GeneSize = (*p).gSize

			nL := len(*ns0) // length of nucleotides

			c0.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
			c1.Sequence = append((*ns1)[:0:0], (*ns1)...) // copy the value

<<<<<<< HEAD
=======
			c0.NDist = append((*nDist0)[:0:0], (*nDist0)...)
			c1.NDist = append((*nDist1)[:0:0], (*nDist1)...)

>>>>>>> iss7
			c0.ErrSequence = make([]byte, nL, nL)
			c1.ErrSequence = make([]byte, nL, nL)

			c0.lSequence = nL
			c1.lSequence = nL

			// handle error
<<<<<<< HEAD
			c0.CheckEM2()
			c1.CheckEM2()

			c0.HandleEM1()
			c1.HandleEM1()

			c0.HandleEM2()
			c1.HandleEM2()
=======
			c0.HandleEM3()
			c1.HandleEM3()
			c0.HandleEM3()
			c1.HandleEM3()
>>>>>>> iss7

			c0.CalFitness()
			c1.CalFitness()

			// generate code name
			c0.GenCode = fmt.Sprintf("%02dC%02d:%04v", g, i, c0.nErr)
			c1.GenCode = fmt.Sprintf("%02dC%02d:%04v", g, i+1, c1.nErr)

			// assign to the populations at index after the fittest
			// population
			(*p).P[p4Size+i] = c0
			(*p).P[p4Size+i+1] = c1

			wg.Done()
		}(i)
	}
	wg.Wait()
}

// New creates new chromosomes from the source sequence
func (p *Population) New(g int) {
	var wg sync.WaitGroup
	// loop through 1/4 for the list
	for i := p34Size; i < pSize; i++ {
		wg.Add(1)

		go func(i int) {
<<<<<<< HEAD
			(*p).P[i] = *newChromo((*p).ns0, (*p).gSize, g, i-p34Size)
=======
			(*p).P[i] = *newChromo((*p).ns0, (*p).nDist, (*p).gSize, g, i-p34Size)
>>>>>>> iss7
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Mutate creates new chromosomes by changing
func (p *Population) Mutate(g int) {
	var wg sync.WaitGroup

	for i := p2Size; i < p34Size; i++ {
		wg.Add(1)
		go func(i int) {
<<<<<<< HEAD
			(*p).P[i].HandleEM2()
			(*p).P[i].CheckEM2()
=======
			(*p).P[i].HandleEM3()
			(*p).P[i].HandleEM3()
			(*p).P[i].HandleEM3()
			(*p).P[i].CalFitness()
>>>>>>> iss7
			(*p).P[i].GenCode = fmt.Sprintf("%02dM%02d:%04v", g, i-p2Size, (*p).P[i].nErr)

			wg.Done()
		}(i)
	}

	wg.Wait()
}

// Sort will sort the data by fitness
func (p *Population) Sort() {
	nc := (*p).P
	sort.Slice(nc[:], func(p, q int) bool {
		return nc[p].Fitness > nc[q].Fitness
	})

	(*p).P = nc
}

// crossover creates new nucleotide sequence by exchanging genes between two
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
