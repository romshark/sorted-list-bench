package rnd

import "math/rand"

func MakeInts(sz int, minVal, maxVal int) []int {
	x := make([]int, sz)
	for i := range x {
		x[i] = RandIntRange(minVal, maxVal)
	}
	return x
}

func RandIntRange(min, max int) int {
	return rand.Intn(max-min) + min
}
