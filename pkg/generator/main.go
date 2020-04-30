package generator

import (
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

const nGeneration = 4

// Start begin the generating process
func Start(tt *models.TimeTable) (*models.TimeTable, error) {
	s0, geneSize, err := (*tt).Decode()
	if err != nil {
		return nil, err
	}

	var p Population
	p.Init(s0, geneSize)

	return nil, nil
}
