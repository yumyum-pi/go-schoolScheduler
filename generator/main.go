package generator

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/models"
)

const n = 32

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
func Init(classes []models.Class, teachers models.Teachers) {
	// make genes
	chrmsm := make(Chromosome, (len(classes) * models.MaxCap))
	var p [n * n]Chromosome = [n * n]Chromosome{}

	// make an array of all the periods
	chrmsm.InitS(&classes)
	totalError := 0

	min := 1000
	max := 0
	///c := 0
	chrmsmLength := len(chrmsm)
	for i := 0; i < n; i++ {
		p[i] = make(Chromosome, chrmsmLength)
		p[i] = chrmsm.InitR()
		p[i].ErrorHandleM1()
		p[i].ErrorHandleM2()

		_, errList := p[i].ErrorCheckM2()
		errNo := len(errList)
		if errNo != 0 {
			totalError += errNo
			if errNo < min {
				min = errNo
			}

			if errNo > max {
				max = errNo
			}
		}
	}

	average := float64(totalError) / float64(n*chrmsmLength) * 100
	fmt.Printf("average=%v, minError=%v, maxError=%v, n=%v chromesomeLength=%v\n", average, min, max, n, len(p[0]))
}
