package lib

// Min returns the minimum of the supplied values.
func Min(vals ...int) int {
	if len(vals) == 0 {
		panic("Zero values")
	}
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum of the supplied values.
func Max(vals ...int) int {
	if len(vals) == 0 {
		panic("Zero values")
	}
	max := vals[0]
	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}
	return max
}