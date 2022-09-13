package myFast

import (
	"math/rand"
	"testing"
)

func bs(data []uint32, target uint32) bool {
	i, j := 0, len(data)
	for i < j {
		mid := (i + j) >> 1
		if data[mid] == target {
			return true
		} else if data[mid] > target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}

func Benchmark_bs(b *testing.B) {
	size := 1<<29 - 1
	if !isFullBSTSize(size) {
		panic("error")
	}
	data := prepareArray(size)
	r := rand.New(rand.NewSource(12))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs(data, r.Uint32())
	}
}
