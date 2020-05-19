package generator

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
	cr "github.com/yumyum-pi/go-schoolScheduler/test/checkerror"
)

func TestChromosome_illegalMutation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc *chromosome // store new chromosome value
	var n byte
	for i := 0; i < iterate; i++ {
		nc = newChromo(ns0, nDist, gSize, 0, i) // create new chromosome

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
	nc.GeneSize = gSize
	for k := 0; k < lk; k++ {
		nc.Sequence = cr.Sequences[k]
		nc.lSequence = len(nc.Sequence)
		nc.ErrSequence = make([]byte, nc.lSequence, nc.lSequence)
		nc.CheckEM1()

		e := checkErrLEqual(cr.Errs[k], nc.ErrSequence, gSize)
		if e != nil {
			t.Error(e)
			return
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

		e = checkErrLEqual(cr.Errs[k], nc.ErrSequence, gSize)
		if e == nil {
			t.Fatalf("was exprecting error but not found at index=%v at k=%v", nc.ErrSequence[i], k)
		}
	}
}

func TestChromosome_CheckEM2(t *testing.T) {
	lk := len(cr.Sequences)
	gSize := 48
	var i int         // length of sequence
	var nc chromosome // store new chromosome value
	nc.GeneSize = gSize
	for k := 0; k < lk; k++ {
		nc.Sequence = cr.Sequences[k]
		nc.lSequence = len(nc.Sequence)
		nc.ErrSequence = make([]byte, nc.lSequence, nc.lSequence)
		nc.CheckEM2()

		e := checkErrLEqual(cr.Errs[k], nc.ErrSequence, gSize)
		if e != nil {
			t.Error(e)
			return
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

		e = checkErrLEqual(cr.Errs[k], nc.ErrSequence, gSize)
		if e == nil {
			t.Fatalf("was exprecting error but not found at index=%v at k=%v", nc.ErrSequence[i], k)
		}
	}
}

func TestChromosome_CheckEM3(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	//var i int // length of sequence
	var e error
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc0 chromosome
	//var nc chromosome
	nc0.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc0.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc0.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc0.ErrSequence = make([]byte, nL, nL)
	nc0.lSequence = nL

	nc1 := nc0
	nc2 := nc0
	for k := 0; k < iterate; k++ {
		shuffleNucleotide(&nc0)
		nc1.Sequence = append(nc0.Sequence[:0:0], nc0.Sequence...)
		nc2.Sequence = append(nc0.Sequence[:0:0], nc0.Sequence...)

		nc1.NDist = append(nc0.NDist[:0:0], nc0.NDist...)
		nc2.NDist = append(nc0.NDist[:0:0], nc0.NDist...)

		nc1.ErrSequence = make([]byte, nL, nL)
		nc2.ErrSequence = make([]byte, nL, nL)

		nc1.CheckEM2()
		nc2.CheckEM3()

		e = checkErrLEqual(nc1.ErrSequence, nc2.ErrSequence, gSize)
		if e != nil {
			t.Error(e)
			return
		}
	}
}

func checkErrLEqual(err0, err1 []byte, gSize int) error {
	l := len(err0)

	if l != len(err1) {
		return fmt.Errorf("length of the err list don't match. err0=%v err1=%v", len(err0), len(err1))
	}
	var n0, n1 byte
	var index int

	errMap0 := make(map[byte]int)
	errMap1 := make(map[byte]int)
	var q0, q1 int
	var ok0, ok1 bool
	// loop through each gene
	for gIndex := 0; gIndex < l; gIndex += gSize {
		for p := 0; p < gSize; p++ {
			index = gIndex + p

			n0 = err0[index]
			n1 = err1[index]

			q0, ok0 = errMap0[n0]
			if !ok0 {
				q0 = 0
			}
			q0++
			errMap0[err0[index]] = q0

			q1, ok1 = errMap1[n1]
			if !ok1 {
				q1 = 0
			}
			q1++
			errMap1[err1[index]] = q1
		}

		for n, q := range errMap0 {
			q1, ok1 = errMap1[n]
			if !ok1 {
				return fmt.Errorf("n=%v at gene=%v not found", n, gIndex)
			}

			if q != q1 {
				return fmt.Errorf("n=%v at gene=%v, q=%v, q1=%v", n, gIndex, q, q1)
			}
		}

	}
	return nil
}

func BenchmarkChromosome_CheckEM1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL
	for i := 0; i < b.N; i++ {
		shuffleNucleotide(&nc)

		nc.CheckEM1()
	}
}

func BenchmarkChromosome_CheckEM2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL
	for i := 0; i < b.N; i++ {
		shuffleNucleotide(&nc)

		nc.CheckEM2()
	}
}

func BenchmarkChromosome_CheckEM3(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL
	for i := 0; i < b.N; i++ {
		shuffleNucleotide(&nc)

		nc.CheckEM3()
	}
}

func TestChromosome_HandleEM1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var e error
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	for k := 0; k < iterate; k++ {
		shuffleNucleotide(&nc)

		nc.CheckEM2()
		nc.HandleEM1()
		nc.CheckEM2()

		e = illegalMutation(ns0, &nc.Sequence, gSize)
		if e != nil {
			t.Error(e)
		}
	}
}

func BenchmarkChromosome_HandleEM1(b *testing.B) {
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	for i := 0; i < b.N; i++ {
		shuffleNucleotide(&nc)

		nc.CheckEM2()
		nc.HandleEM1()
		nc.CheckEM2()
	}
}

func TestChromosome_HandleEM2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var e error
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	var nErr0, nErr1, nErr2 int
	for k := 0; k < iterate; k++ {
		shuffleNucleotide(&nc)

		nc.CheckEM2()
		nErr0 = nc.nErr
		nc.HandleEM1()
		nc.CheckEM2()
		nErr1 = nc.nErr
		if nErr0/2 <= nErr1 {
			t.Errorf("nErr0=%v\tnErr1=%v\n", nErr0, nErr1)
		}
		nc.HandleEM2()
		nc.CheckEM2()
		nErr2 = nc.nErr
		if nErr1/2 <= nErr2 {
			t.Errorf("nErr0=%v\tnErr1=%v\n", nErr1, nErr2)
		}
		e = illegalMutation(ns0, &nc.Sequence, gSize)
		if e != nil {
			t.Error(e)
		}
	}
}

func BenchmarkChromosome_HandleEM2(b *testing.B) {
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	for i := 0; i < b.N; i++ {
		shuffleNucleotide(&nc)
		nc.CheckEM2()
		nc.HandleEM1()
		nc.HandleEM2()
	}
}

func TestChromosome_HandleEM3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var e error
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	//var nErr0, nErr1, nErr2 int
	for k := 0; k < iterate; k++ {
		shuffleNucleotide(&nc)
		nc.HandleEM3()

		e = illegalMutation(ns0, &nc.Sequence, gSize)
		if e != nil {
			t.Error(e)
		}
	}
}

func BenchmarkChromosome_HandleEM3(b *testing.B) {
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	for i := 0; i < b.N; i++ {
		shuffleNucleotide(&nc)
		nc.HandleEM3()
	}
}

func TestChromosome_CalFitness(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL

	for k := 0; k < iterate; k++ {
		shuffleNucleotide(&nc)
		nc.HandleEM3()
		nc.CalFitness()
	}
}

func BenchmarkChromosome_CalFitness(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	for i := 0; i < b.N; i++ {
		// decode the pkgs to ns0 and gene-size
		newChromo(ns0, nDist, gSize, 0, i) // create new chromosome
	}
}

func TestChromosome_SwapNucleotide(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	var nc chromosome
	nc.GeneSize = gSize
	nL := len(*ns0) // length of nucleotides

	nc.Sequence = append((*ns0)[:0:0], (*ns0)...) // copy the value
	nc.NDist = append((*nDist)[:0:0], (*nDist)...)
	nc.ErrSequence = make([]byte, nL, nL)
	nc.lSequence = nL
	for k := 0; k < 64; k++ {
		shuffleNucleotide(&nc)
		for i, n := range nc.Sequence {
			p := i % nc.GeneSize
			j := nMatch(&nc.Sequence, gSize, i)
			distI := (int((n - 1)) * gSize) + p
			k := nc.NDist[distI]
			if int(k) != j {
				t.Errorf("no match found k=%v j =%v", k, j)
			}
		}
	}
}
