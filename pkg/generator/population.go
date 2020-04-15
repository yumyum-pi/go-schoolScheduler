package generator

import (
	"fmt"
	"sort"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

const populationSize = 32
const populationHalfSize = populationSize / 2
const populationQuaterSize = populationSize / 4

// ByEF is a sorting struct by error and fitness
type ByEF []Chromosome

func (bf ByEF) Len() int      { return len(bf) }
func (bf ByEF) Swap(i, j int) { bf[i], bf[j] = bf[j], bf[i] }
func (bf ByEF) Less(i, j int) bool {
	// if error in not 0
	if bf[i].ErrNo != 0 && bf[j].ErrNo != 0 {
		return bf[i].ErrNo < bf[j].ErrNo
	}
	if bf[i].ErrNo != 0 && bf[j].ErrNo == 0 {
		return false
	}
	if bf[i].ErrNo == 0 && bf[j].ErrNo != 0 {
		return true
	}
	return bf[i].Fitness < bf[j].Fitness
}

// Population is a fixed arry of size 64 containing chromosomes
type Population struct {
	P            [populationSize]Chromosome
	totalError   int
	minErr       int
	maxErr       int
	errChrmsmIDs []int
	class        *[]models.Class
}

// Init create the initail population
func (p *Population) Init(classes *[]models.Class) {
	(*p).class = classes
	// make a new chromosome var
	var chrmsm Chromosome
	chrmsm.Genes = make([]models.Period, (len(*classes) * models.MaxCap))

	// make an array of all the periods
	chrmsm.InitS(classes)

	for i := 0; i < populationSize; i++ {
		p.P[i] = chrmsm.InitR()
		p.P[i].ErrorHandle()
		p.P[i].ErrorCheck()
		p.P[i].CalFitness(false)

		errNo := len(p.P[i].ErrIDs)
		//fmt.Println(i, errNo)
		if errNo != 0 {
			p.errChrmsmIDs = append(p.errChrmsmIDs, i)
			p.totalError += errNo
			if errNo < p.minErr {
				p.maxErr = errNo
			}

			if errNo > p.maxErr {
				p.maxErr = errNo
			}
		} else {
			p.minErr = 0
		}
	}
}

// Print will write data to console
func (p *Population) Print() {

	//average := float64((*p).totalError) / float64(populationSize*(*p).P[0].Length()) * 100
	fmt.Printf("nErrChrmsm=%v, minError=%v, maxError=%v, n=%v chromesomeLength=%v\n", len((*p).errChrmsmIDs), (*p).minErr, (*p).maxErr, populationSize, (*p).P[0].Length())
}

// PrintChromo will write data to console
func (p *Population) PrintChromo() {

	for index, chrmsm := range (*p).P {
		fmt.Printf("i=%v\terrN=%v\terrL=%v\tfitness=%v\n", index, chrmsm.ErrNo, len(chrmsm.ErrIDs), chrmsm.Fitness)

	}
}

// Next create's the next gene of chromosome
func (p *Population) Next() {
	(*p).Wip()
	(*p).CrossOver()
	(*p).Sort()
}

// CrossOver creates new geners form the existing genes
func (p *Population) CrossOver() {
	// make crossOver of chromosome 1 and 2

	for p1Index := 0; p1Index < populationHalfSize; p1Index += 2 {
		p2Index := p1Index + 1

		var chrmsm1, chrmsm2 Chromosome

		chrmsm1.Genes, chrmsm2.Genes = crossOver(p.P[p1Index].Genes, p.P[p2Index].Genes)

		chrmsm1.ErrorHandle()
		chrmsm2.ErrorHandle()
		chrmsm1.ErrorCheck()
		chrmsm2.ErrorCheck()
		//	fmt.Println("> ErrorCheck M3=", chrmsm1.ErrorCheckM3((*p).class))
		//	fmt.Println("> ErrorCheck M3=", chrmsm2.ErrorCheckM3((*p).class))
		// chrmsm2.ErrorHandle()
		chrmsm1.CalFitness(false)
		chrmsm2.CalFitness(false)
		//fmt.Printf("chrmsm1 err=%v fitness=%v\n", chrmsm1.ErrNo, chrmsm1.Fitness)
		//fmt.Printf("chrmsm2 err=%v fitness=%v\n", chrmsm2.ErrNo, chrmsm2.Fitness)

		// the chromosome to the population
		(*p).P[p1Index+populationHalfSize] = chrmsm1
		(*p).P[p2Index+populationHalfSize] = chrmsm2
	}
	p.PrintChromo()

}

// Wip will delete
func (p *Population) Wip() {
	nc := make([]Chromosome, populationSize)
	copy(nc, (*p).P[:populationHalfSize])
	copy((*p).P[:], nc)

}

// Sort will sort the data by fitness
func (p *Population) Sort() {
	c := (*p).P[:]
	// nc := make([]Chromosome, populationSize)
	// sort by fitness
	sort.Sort(ByEF(c))
	copy((*p).P[:], c)

}

func crossOver(p1, p2 []models.Period) ([]models.Period, []models.Period) {
	var gene1, gene2 []models.Period
	l1 := len(p1)
	l2 := len(p2)
	if l1 != l2 {
		fmt.Println("Kill me")
	}
	classIndex := 1
	for geneIndex := 0; geneIndex < l1; geneIndex += models.MaxCap {
		var newClassGenes1 []models.Period
		var newClassGenes2 []models.Period
		// if class is even
		if classIndex%2 == 0 {
			newClassGenes1 = p1[geneIndex : geneIndex+models.MaxCap]
			newClassGenes2 = p2[geneIndex : geneIndex+models.MaxCap]
		} else {
			newClassGenes1 = p2[geneIndex : geneIndex+models.MaxCap]
			newClassGenes2 = p1[geneIndex : geneIndex+models.MaxCap]
		}
		classIndex++
		gene1 = append(gene1, newClassGenes1...)
		gene2 = append(gene2, newClassGenes2...)
	}

	if len(p2) != len(gene1) {
		fmt.Println("length not equal. p2=", len(p2), "\tgene1=", len(gene1))
	}
	if len(p1) != len(gene2) {
		fmt.Println("length not equal. p1=", len(p1), "\tgene2=", len(gene2))
	}
	g1 := make([]models.Period, l1)
	g2 := make([]models.Period, l2)
	copy(g1, gene1)
	copy(g2, gene2)

	return g1, g2
}

/*
 */
