package utils

import (
	"testing"
)

var is = []int{
	220,
	2020,
	1<<16 - 1,
	1<<24 - 1,
}
var bs = []B256{
	{0, 220},
	{7, 228},
	{255, 255},
	{255, 255, 255},
}

// TODO add unit test
func TestB256_Encode(t *testing.T) {
	for _, i := range is {
		b := make(B256, 2)
		b.Encode(i)
		t.Errorf("%v", b)
	}
}

// TODO add unit test
func TestB256_Decode(t *testing.T) {
	for i, b := range bs {
		if b.Decode() != is[i] {
			t.Errorf(">Error : was expecting %v but got %v", is[i], bs[i].Decode())
		}
	}

}
