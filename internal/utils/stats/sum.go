package stats

// Sum return the sun of the given int slice
func Sum(i []int) (sum int) {
	for _, s := range i {
		sum += s
	}
	return
}

// SumF32 return the sun of the given float32 slice
func SumF32(i []float32) (sum float32) {
	for _, s := range i {
		sum += s
	}
	return
}

// Mean return the means of the given int slice
func Mean(i []int) (mean float64) {
	sum := Sum(i)
	mean = float64(sum) / float64(len(i))
	return
}

// Variance return the varient of the given int slice
func Variance(i []int) (mean float64) {
	m := Mean(i)
	var x2Sum float64
	for _, x := range i {
		x2Sum += float64(x * x)
	}
	m2 := m * m
	v := (x2Sum / float64(len(i))) - m2
	return v
}

// VarianceInt return the varient of the given int slice
func VarianceInt(i []int) (mean int) {
	m := Mean(i)
	var x2Sum float64
	for _, x := range i {
		x2Sum += float64(x * x)
	}
	m2 := m * m
	v := (x2Sum / float64(len(i))) - m2
	return int(v)
}
