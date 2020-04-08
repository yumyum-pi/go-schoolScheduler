package stats

import "testing"

var tIntSlice []int = []int{600, 470, 170, 430, 300}
var sum int = 1970
var mean float64 = 394
var variance float64 = 21704

func TestSum(t *testing.T) {
	s := Sum(tIntSlice)
	if s != sum {
		t.Errorf("> Error: expected sum=%v, got sum=%v", sum, s)
	}
}

func TestMean(t *testing.T) {
	m := Mean(tIntSlice)
	if m != mean {
		t.Errorf("> Error: expected mean=%v, got mean=%v", mean, m)
	}
}

func TestVariance(t *testing.T) {
	v := Variance(tIntSlice)
	if v != variance {
		t.Errorf("> Error: expected Variance=%v, got Variance=%v", variance, v)
	}
}
