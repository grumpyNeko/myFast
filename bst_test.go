package myFast

import (
	gen "github.com/grumpyNeko/myFuzz"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_basic(t *testing.T) {
	bst := arrayToBST([]uint32{0, 1, 2, 3, 4, 5, 6})
	for i := 5; i < 10; i++ {
		println(getFromBst(bst, uint32(i)))
	}
	println(bst)
}

func Benchmark_bst(b *testing.B) {
	size := 1<<29 - 1
	data := prepareArray(size)
	bst := arrayToBST(data)
	r := rand.New(rand.NewSource(12))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getFromBst(bst, r.Uint32())
		//for i:=0; i<len(data); i++ {
		//	getFromBst(bst, data[i])
		//}
	}
}

func Test_pp(t *testing.T) {
	size := 1<<29 - 1
	data := prepareArray(size)
	bst := arrayToBST(data)

	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < len(data); i++ {
		getFromBst(bst, data[i])
	}
}

func Test_random(t *testing.T) {
	size := 1<<29 - 1
	g := gen.NewUint64Generator(0, 1<<32, nil, true, 1<<29-1, false)
	data := g.Generate() //prepareArray(size)
	// todo: 32bit
	bst := arrayToBST(data)

	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < len(data); i++ {
		getFromBst(bst, data[i])
	}
}
