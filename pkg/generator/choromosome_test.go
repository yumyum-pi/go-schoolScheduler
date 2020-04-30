package generator

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
)

func TestChromosome_illegalMutation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	pkgs := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := pkgs.Decode()

	var nc *chromosome // store new chromosome value

	for i := 0; i < 256; i++ {
		nc = newChromo(ns0, gSize) // create new chromosome

		if e := illegalMutation(ns0, &(nc.Sequence), gSize); e != nil {
			t.Error(e)
		}

		// make illegal mutation
		// type of nucleotide not found
		nc = newChromo(ns0, gSize) // create new chromosome
		n := byte(rand.Intn(255))  // create random byte

		for n == nc.Sequence[n] {
			n = byte(rand.Intn(255)) // create random byte
		}
		nc.Sequence[n] = n
		if e := illegalMutation(ns0, &(nc.Sequence), gSize); e == nil {
			t.Errorf("was exprecting an error but not found")
		}
	}
}
