package requestlist

import (
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// TClassRL is test requrest list for classes
var TClassRL = Class{
	models.TClassIDL[0].Bytes(): 9,
	models.TClassIDL[1].Bytes(): 10,
	models.TClassIDL[2].Bytes(): 4,
	models.TClassIDL[3].Bytes(): 3,
	models.TClassIDL[4].Bytes(): 8,
	models.TClassIDL[5].Bytes(): 12,
	models.TClassIDL[6].Bytes(): 2,
	models.TClassIDL[7].Bytes(): 5,
	models.TClassIDL[8].Bytes(): 1,
	models.TClassIDL[9].Bytes(): 6,
}

var getCRL = func(j, l int) Class {
	crl := make(Class)
	for i := 0; i < l; i++ {
		crl[models.TClassL[9-i].ID.Bytes()] = models.TSubjectL[j].Req
	}
	return crl
}

// TSRL is test request list for subjects
var TSRL = Subject{
	models.TSubjectIDL[0].Type: getCRL(0, 10),
	models.TSubjectIDL[1].Type: getCRL(1, 9),
	models.TSubjectIDL[2].Type: getCRL(2, 8),
	models.TSubjectIDL[3].Type: getCRL(3, 7),
	models.TSubjectIDL[4].Type: getCRL(4, 6),
	models.TSubjectIDL[5].Type: getCRL(5, 5),
	models.TSubjectIDL[6].Type: getCRL(6, 4),
	models.TSubjectIDL[7].Type: getCRL(7, 3),
	models.TSubjectIDL[8].Type: getCRL(8, 2),
	models.TSubjectIDL[9].Type: getCRL(9, 1),
}
