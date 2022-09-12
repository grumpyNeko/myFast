package myFast

import (
	"math/rand"
	"testing"
)

// 1>>29 = 512M // // 2GB
func prepareArray(cap int) []uint32 {
	r := rand.New(rand.NewSource(1))
	keys := make([]uint32, cap, cap)
	for i:=0; i<cap; i++ {
		keys[i] = r.Uint32()
	}
	return keys
}

func arrayToBST(data []uint32) []uint32 {
	size := len(data)
	if !isFullBSTSize(size) {
		panic("given size cannot apply to a full bst")
	}
	tree := make([]uint32, size, size)

	setNode(0, 0, size, data, tree)
	return tree
}

func isFullBSTSize(size int) bool {
	for i:=0; i<33; i++ {
		if size == 1 << i - 1 {
			return true
		}
	}
	return false
}

func Test_arrayToBST(t *testing.T) {
	bst := arrayToBST([]uint32{0, 1, 2, 3, 4, 5, 6})
	println(bst)
}

func setNode(nodePos int, low int, high int, data []uint32, tree []uint32) {
	if low > high {
		return
	}
	mid := (low+high) / 2
	tree[nodePos] = data[mid]
	l := left(nodePos)
	if l < len(tree) {
		setNode(l, low, mid-1, data, tree)
	}
	r := right(nodePos)
	if r < len(tree) {
		setNode(r, mid+1, high, data, tree)
	}
}

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

type queue struct {
	data []int
	head  int
	tail  int
}

func NewQueue(cap int) *queue {
	return &queue{
		data: make([]int, cap, cap),
		head: 0,
		tail: 0,
	}
}

func (q *queue)push_back(ele int) {
	q.data[q.tail] = ele
	q.tail ++
}

func (q *queue)not_empty() bool {
	return q.head < q.tail
}

func (q *queue)pop_front() int {
	q.head ++
	return q.data[q.head-1]
}

// todo: 数组表示的BST, 天然就是levelOrder
func levelOrder(tree []uint32) []uint32 {
	cap := len(tree)
	q := NewQueue(cap)
	root := 0
	re := make([]uint32, 0, cap)
	q.push_back(root)
	for q.not_empty() {
		e := q.pop_front()
		l := left(e)
		r := right(e)
		if l < cap {
			q.push_back(l)
		}
		if r < cap {
			q.push_back(r)
		}
		re = append(re, tree[e])
	}
	return re
}

//func Test_levelOrder(t *testing.T) {
//	cap := 7
//	re := levelOrder(prepareTree(cap))
//	println(re)
//}

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

func Benchmark_bs(b *testing.B){
	data := prepareArray(1>>29)
	r := rand.New(rand.NewSource(12))
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bs(data, r.Uint32())
	}
}
