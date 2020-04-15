package generator

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

const nGeneration = 4

func chrmsmByteMatch(c1, c2 []byte) bool {
	// If one is nil, the other must also be nil.
	if (c1 == nil) != (c2 == nil) {
		return false
	}

	if len(c1) != len(c2) {
		return false
	}

	for i := range c1 {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}

// Init starts the process
func Init(classes []models.Class, teachers models.Teachers) Chromosome {

	var p Population

	p.Init(&classes)
	p.Sort()
	p.PrintChromo()
	p.Print()
	fmt.Println("Init population")

	for generationIndex := 0; generationIndex < nGeneration; generationIndex++ {
		fmt.Println("population generation:", generationIndex)
		p.Next()
		p.Print()
	}

	// choose the best chromosome
	return p.P[0]
}
