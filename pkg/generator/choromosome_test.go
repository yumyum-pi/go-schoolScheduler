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
	lk := len(cr.Sequences)
	gSize := 48
	var i int         // length of sequence
	var nc chromosome // store new chromosome value
	for k := 0; k < lk; k++ {
		nc = *newChromo(&cr.Sequences[k], gSize)
		nc.CheckEM1()

		e := checkErrLEqual(cr.Errs[k], nc.ErrSequence)
		if e != nil {
			t.Fatal(e)
		}

		// manipulate data to give error
		i = len(nc.Sequence) - 1
		for i >= 0 {
			if nc.ErrSequence[i] != 0 {
				nc.ErrSequence[i] = 0
				break
			}
			i--
		}

		e = checkErrLEqual(cr.Errs[k], nc.ErrSequence)
		if e == nil {
			t.Fatalf("was exprecting error but not found at index=%v at k=%v", nc.ErrSequence[i], k)
		}
		break
		/**/
	}
}

/*
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
*/
func checkErrLEqual(err0, err1 []byte) error {
	l := len(err0)

	if l != len(err1) {
		return fmt.Errorf("length of the err list don't match. err0=%v err1=%v", len(err0), len(err1))
	}
	var n0, n1 byte
	// loop through each gene
	for sIndex := 0; sIndex < l; sIndex++ {
		n0 = err0[sIndex]
		n1 = err1[sIndex]
		if n0 != n1 {
			return fmt.Errorf("nucleotides at sIndex=%v are different. n0=%v, n1=%v", sIndex, n0, n1)
		}
	}
	return nil
}

/*
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
	var gSize, err0 int
	var e error
	// get information from the file
	pkgs := file.ReadRand(inputDir)
	for k := 0; k < 1; k++ {
		var nc *chromosome

		// decode the pkgs to ns0 and gene-size
		ns0, gSize, _ = pkgs.Decode()
		nc = newChromo(ns0, gSize) // create new chromosome

		nc.CheckEM2()
		nc.Print()
		nc.HandleEM1()
		nc.Print()
		err0 = len(nc.ErrIndexL)
		nc.CheckEM2()
		nc.Print()
		if err0 != len(nc.ErrIndexL) {
			t.Errorf("false error resolve from handleEM1, err0=%v, err=%v", err0, len(nc.ErrIndexL))
		}
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
	var gSize, err0 int
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
		err0 = len(nc.ErrIndexL)
		nc.CheckEM2()
		if err0 != len(nc.ErrIndexL) {
			t.Errorf("false error resolve from handleEM2")
		}
		//nc.Print()
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

func TestChromosome_CheckSafeSwap(t *testing.T) {

}
*/
