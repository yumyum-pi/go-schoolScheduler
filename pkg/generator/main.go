package generator

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

const nGeneration = 4

var tpCount map[byte]int = make(map[byte]int) // total period counter of a teacher

// Start begin the generating process
func Start(tt *models.TimeTable) (*models.TimeTable, error) {
	b, max := (*tt).Decode()
	bl := len(b)
	var tID byte // teacher ID
	var tp int   // total periods of a teacher
	var ok bool  // if the teacher is persent in the map
	for i := 0; i < bl; i += max {
		for j := 0; j < max; j++ {
			tID = b[(i + j)]

			tp, ok = tpCount[tID]
			if !ok {
				tp = 0
			}
			tp++
			tpCount[tID] = tp
		}

	}

	for _, tp := range tpCount {
		if tp >= max {
			return nil, fmt.Errorf("> Error: Data received is invalid")
		}
	}
	var p Population
	p.Init(&b, max)

	return nil, nil
}

/*
// Init starts the process
func Init(tt *models.TimeTable) Chromosome {

	var p Population

	p.Init(&tt)
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
*/
