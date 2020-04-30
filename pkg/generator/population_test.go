package generator

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
)

var inputDir = "/Users/vivekrawat/go/src/github.com/yumyum-pi/go-schoolScheduler/test/inputs"

func TestPopulation_newChromo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	pkgs := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize := pkgs.Decode()

	// create new chromosome
	nc := newChromo(&ns0, gSize)

	if e := illegalMutation(&ns0, &(nc.Nucleotides), gSize); e != nil {
		t.Error(e)
	}
}
