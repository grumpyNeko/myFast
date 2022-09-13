package myFast

import "math/rand"

// 1<<29是512M // // 2GB
func prepareArray(cap int) []uint32 {
	r := rand.New(rand.NewSource(1))
	keys := make([]uint32, cap, cap)
	for i := 0; i < cap; i++ {
		keys[i] = r.Uint32()
	}
	return keys
}
