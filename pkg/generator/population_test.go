package generator

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

var inputDir = "/Users/vivekrawat/go/src/github.com/yumyum-pi/go-schoolScheduler/test/inputs/"

func TestNewChromo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	// create new chromosome
	nc := newChromo(ns0, nDist, gSize, 0, 0)

	if e := illegalMutation(ns0, &(nc.Sequence), gSize); e != nil {
		t.Error(e)
	}
}

func BenchmarkPopulation_Init(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	p := CreatePopulation(ns0, nDist, gSize, int(req.NNType))

	for i := 0; i < b.N; i++ {
		p.Init()
	}
}

func TestPopulation_Sort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	p := CreatePopulation(ns0, nDist, gSize, int(req.NNType))
	p.Init()

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
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	// create new chromosome
	c0 := newChromo(ns0, nDist, gSize, 0, 0)
	c1 := newChromo(ns0, nDist, gSize, 0, 0)

	n0, n1 := crossOver(&c0.Sequence, &c1.Sequence, gSize)

	if e := illegalMutation(ns0, n0, gSize); e != nil {
		t.Error(e)
	}

	if e := illegalMutation(ns0, n1, gSize); e != nil {
		t.Error(e)
	}
}

func BenchmarkCrossOver(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	// create new chromosome
	c0 := newChromo(ns0, nDist, gSize, 0, 0)
	c1 := newChromo(ns0, nDist, gSize, 0, 0)
	for i := 0; i < b.N; i++ {
		crossOver(&c0.Sequence, &c1.Sequence, gSize)
	}
}

func TestPopulation_CrossOver(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	p := CreatePopulation(ns0, nDist, gSize, int(req.NNType))
	p.Init()
	p.CrossOver(1)
}

func TestPopulation_Mutate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	p := CreatePopulation(ns0, nDist, gSize, int(req.NNType))
	p.Init()
	p.CrossOver(1)
	p.Mutate(1)
}

func TestPopulation_New(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	p := CreatePopulation(ns0, nDist, gSize, int(req.NNType))
	p.Init()
	p.CrossOver(1)
	p.Mutate(1)
	p.New(1)
}

func BenchmarkPopulation_Next(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	// get information from the file
	req := file.ReadRand(inputDir)

	// decode the pkgs to ns0 and gene-size
	ns0, gSize, _ := models.Decode(&req.Pkgs, req.GSize)
	nDist := nDistribution(ns0, gSize, int(req.NNType))
	p := CreatePopulation(ns0, nDist, gSize, int(req.NNType))
	p.Init()
	for i := 0; i < b.N; i++ {
		p.Next(i)
	}
}
