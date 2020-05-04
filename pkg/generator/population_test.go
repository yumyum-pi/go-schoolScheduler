package generator

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
)

var inputDir = "/Users/vivekrawat/go/src/github.com/yumyum-pi/go-schoolScheduler/test/inputs"

func TestNewChromo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	pkgs := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := pkgs.Decode()

	// create new chromosome
	nc := newChromo(ns0, gSize, 0, 0)

	if e := illegalMutation(ns0, &(nc.Sequence), gSize); e != nil {
		t.Error(e)
	}
}

func BenchmarkPopulation_Init(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	pkgs := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := pkgs.Decode()

	var p Population

	for i := 0; i < b.N; i++ {
		p.Init(ns0, gSize)
	}
}

func TestPopulation_Sort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	pkgs := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := pkgs.Decode()

	var p Population
	p.Init(ns0, gSize)

	fitness := make([]int, pSize)

	for i, c := range p.P {
		fitness[pSize-i-1] = c.Fitness
	}

	if !sort.IntsAreSorted(fitness) {
		fmt.Println(fitness)
		t.Errorf("chromosome is not sorted")
	}
}

func TestCrossOver(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	pkgs := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := pkgs.Decode()

	// create new chromosome
	c0 := newChromo(ns0, gSize, 0, 0)
	c1 := newChromo(ns0, gSize, 0, 0)

	c0.CheckEM2()
	c0.HandleEM1()
	c0.HandleEM2()

	c1.CheckEM2()
	c1.HandleEM1()
	c1.HandleEM2()

	n0, n1 := crossOver(&c0.Sequence, &c1.Sequence, gSize)

	if e := illegalMutation(ns0, n0, gSize); e != nil {
		t.Error(e)
	}

	if e := illegalMutation(ns0, n1, gSize); e != nil {
		t.Error(e)
	}
}
