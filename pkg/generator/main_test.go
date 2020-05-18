package generator

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// ittrate no
var ittrate int = 1

func TestNDistribution(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)

	fmt.Println(req.NNType)
	//var nc *chromosome // store new chromosome value
	//var n byte
	for i := 0; i < ittrate; i++ {
		nDist := nDistrubution(ns0, gSize, int(req.NNType))
		printNDistribution(nDist, gSize)
		for i, n := range *ns0 {
			p := i % gSize
			j := nMatch(ns0, gSize, i)
			distI := (int((n - 1)) * gSize) + p
			k := (*nDist)[distI]
			if int(k) != j {
				t.Errorf("no match found k=%v j =%v", k, j)
			}
		}
	}
}

// MatchN check nucleotides at the given position in each gene
// and returns the index of the matching nucleotide . If no match is
// found then -1 is returned
func nMatch(c *[]byte, gSize int, sIndex0 int) (y int) {
	n := (*c)[sIndex0]
	l := len(*c)
	// calculate gene position
	p := sIndex0 % gSize

	// loop though each gene
	for gIndex := 0; gIndex < l; gIndex += gSize {
		sIndex1 := gIndex + p // calc the sequence index

		// check for matching nucleotide
		// avoid same sequence index
		if (*c)[sIndex1] == n {
			y++
		}
	}
	return
}
