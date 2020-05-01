package generator

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
	cr "github.com/yumyum-pi/go-schoolScheduler/test/checkerror"
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

func TestChromosome_CheckEM1(t *testing.T) {
	l := len(cr.Sequences)
	gSize := 48
	for k := 0; k < l; k++ {

		var nc chromosome // store new chromosome value
		nc.GeneSize = gSize
		nc.Sequence = cr.Sequences[k]

		nc.CheckEM1()

		e := checkErrLEqual(cr.Errs[k], nc.ErrIndexL)
		if e != nil {
			t.Fatal(e)
		}

		// manipulate data to give error
		r := rand.Intn(len(nc.ErrIndexL))
		nc.ErrIndexL[r] = 0
		e = checkErrLEqual(cr.Errs[k], nc.ErrIndexL)
		if e == nil {
			t.Fatalf("was exprecting error but not found")
		}
	}
}

func checkErrLEqual(err1, err2 []int) error {
	mapErr2 := make(map[int]bool)
	// make the map
	for _, sIndex := range err2 {
		mapErr2[sIndex] = true
	}

	// compare mapErr2 to err1
	for _, sIndex := range err1 {
		_, ok := mapErr2[sIndex]
		if !ok {
			return fmt.Errorf("%v not found in err2", sIndex)
		}
	}
	return nil
}
