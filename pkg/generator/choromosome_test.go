package generator

import (
	"fmt"
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
	ns0, gSize := pkgs.Decode()

	var nc *chromosome // store new chromosome value

	for i := 0; i < 100; i++ {
		nc = newChromo(&ns0, gSize) // create new chromosome

		if e := illegalMutation(&ns0, &(nc.Nucleotides), gSize); e != nil {
			t.Error(e)
		}

		// make illegal mutation
		// type of nucleotide not found
		nc = newChromo(&ns0, gSize) // create new chromosome
		n := byte(rand.Intn(255))   // create random byte

		// check the value of the n in the new sequence
		if n != nc.Nucleotides[n] {
			fmt.Println(nc.Nucleotides[n])
			nc.Nucleotides[n] = n
		}
		if e := illegalMutation(&ns0, &(nc.Nucleotides), gSize); e == nil {
			t.Errorf("was exprecting an error but not found")
		}
	}
}
