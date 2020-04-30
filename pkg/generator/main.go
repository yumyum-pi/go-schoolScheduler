package generator

import (
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

const nGeneration = 4

// Start begin the generating process
func Start(s0 *[]byte, geneSize int) (*models.TimeTable, error) {
	var p Population
	p.Init(s0, geneSize)

	return nil, nil
}
