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
	models.TSubjectIDL[0].Bytes(): getCRL(0, 10),
	models.TSubjectIDL[1].Bytes(): getCRL(1, 9),
	models.TSubjectIDL[2].Bytes(): getCRL(2, 8),
	models.TSubjectIDL[3].Bytes(): getCRL(3, 7),
	models.TSubjectIDL[4].Bytes(): getCRL(4, 6),
	models.TSubjectIDL[5].Bytes(): getCRL(5, 5),
	models.TSubjectIDL[6].Bytes(): getCRL(6, 4),
	models.TSubjectIDL[7].Bytes(): getCRL(7, 3),
	models.TSubjectIDL[8].Bytes(): getCRL(8, 2),
	models.TSubjectIDL[9].Bytes(): getCRL(9, 1),
}
