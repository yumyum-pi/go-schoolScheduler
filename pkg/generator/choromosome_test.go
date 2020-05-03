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
	var n byte
	for i := 0; i < 256; i++ {
		nc = newChromo(ns0, gSize) // create new chromosome

		if e := illegalMutation(ns0, &(nc.Sequence), gSize); e != nil {
			t.Error(e)
		}
		n = byte(rand.Intn(255)) // create random byte

		for n == nc.Sequence[n] {
			n = byte(rand.Intn(255)) // create random byte
		}
		nc.Sequence[n] = n
		if e := illegalMutation(ns0, &(nc.Sequence), gSize); e == nil {
			t.Error(e)
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
		nc.ErrIndexL[0] = -1
		e = checkErrLEqual(cr.Errs[k], nc.ErrIndexL)
		if e == nil {
			t.Fatalf("was exprecting error but not found at index=%v at k=%v", nc.ErrIndexL[0], k)
		}
	}
}

func TestChromosome_CheckEM2(t *testing.T) {
	l := len(cr.Sequences)
	gSize := 48
	for k := 0; k < l; k++ {

		var nc chromosome // store new chromosome value
		nc.GeneSize = gSize
		nc.Sequence = cr.Sequences[k]

		nc.CheckEM2()

		e := checkErrLEqual(cr.Errs[k], nc.ErrIndexL)
		if e != nil {
			t.Fatal(e)
		}
		// manipulate data to give error
		nc.ErrIndexL[0] = -1
		e = checkErrLEqual(cr.Errs[k], nc.ErrIndexL)
		if e == nil {
			t.Fatalf("was exprecting error but not found at index=%v at k=%v", nc.ErrIndexL[0], k)
		}
	}
}

func checkErrLEqual(err1, err2 []int) error {
	// map for all the index in err2
	mapErr2 := make(map[int]bool)
	if len(err1) != len(err2) {
		return fmt.Errorf("length of the err list don't match. err1=%v err2=%v", len(err1), len(err2))
	}
	// add the values to the map
	for i, sIndex := range err2 {
		_, ok := mapErr2[sIndex]
		if ok {
			return fmt.Errorf("map should not have repeating sIndex=%v in err2 at i=%v", sIndex, i)
		}
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

func BenchmarkChromosome_CheckEM1(b *testing.B) {
	l := len(cr.Sequences)
	gSize := 48

	for i := 0; i < b.N; i++ {
		for k := 0; k < l; k++ {
			var nc chromosome // store new chromosome value
			nc.GeneSize = gSize
			nc.Sequence = cr.Sequences[k]

			nc.CheckEM1()
		}
	}
}

func BenchmarkChromosome_CheckEM2(b *testing.B) {
	l := len(cr.Sequences)
	gSize := 48

	for i := 0; i < b.N; i++ {
		for k := 0; k < l; k++ {
			var nc chromosome // store new chromosome value
			nc.GeneSize = gSize
			nc.Sequence = cr.Sequences[k]

			nc.CheckEM2()
		}
	}
}

func TestChromosome_HandleEM1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ns0 *[]byte
	var gSize int
	var e error
	// get information from the file
	pkgs := file.ReadRand(inputDir)
	for k := 0; k < 256; k++ {
		var nc *chromosome

		// decode the pkgs to ns0 and gene-size
		ns0, gSize, _ = pkgs.Decode()
		nc = newChromo(ns0, gSize) // create new chromosome

		nc.CheckEM2()
		nc.HandleEM1()
		e = illegalMutation(ns0, &nc.Sequence, gSize)
		if e != nil {
			t.Error(e)
		}
	}
}

func BenchmarkChromosome_HandleEM1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var ns0 *[]byte
	var gSize int
	//	var e error

	var nc chromosome // store new chromosome value
	nc.GeneSize = gSize

	// get information from the file
	pkgs := file.ReadRand(inputDir)
	for i := 0; i < b.N; i++ {
		var nc *chromosome
		// decode the pkgs to ns0 and gene-size
		ns0, gSize, _ = pkgs.Decode()
		nc = newChromo(ns0, gSize) // create new chromosome

		nc.CheckEM2()
		nc.HandleEM1()

	}
}

func TestChromosome_HandleEM2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ns0 *[]byte
	var gSize int
	var e error
	// get information from the file
	pkgs := file.ReadRand(inputDir)
	for k := 0; k < 256; k++ {
		var nc *chromosome

		// decode the pkgs to ns0 and gene-size
		ns0, gSize, _ = pkgs.Decode()
		nc = newChromo(ns0, gSize) // create new chromosome

		nc.CheckEM2()
		nc.HandleEM1()
		e = nc.HandleEM2()
		if e != nil {
			t.Error(e)
			continue
		}
		nc.Print()
		e = illegalMutation(ns0, &nc.Sequence, gSize)
		if e != nil {
			t.Error(e)
		}
	}
}

func BenchmarkChromosome_HandleEM2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var ns0 *[]byte
	var gSize int
	//	var e error

	var nc chromosome // store new chromosome value
	nc.GeneSize = gSize

	// get information from the file
	pkgs := file.ReadRand(inputDir)
	ns0, gSize, _ = pkgs.Decode()
	for i := 0; i < b.N; i++ {
		var nc *chromosome
		// decode the pkgs to ns0 and gene-size
		nc = newChromo(ns0, gSize) // create new chromosome

		nc.CheckEM2()
		nc.HandleEM1()
		nc.HandleEM2()
	}
}
