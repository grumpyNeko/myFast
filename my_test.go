package myFast

type queue struct {
	data []int
	head int
	tail int
}

func NewQueue(cap int) *queue {
	return &queue{
		data: make([]int, cap, cap),
		head: 0,
		tail: 0,
	}
}

func (q *queue) push_back(ele int) {
	q.data[q.tail] = ele
	q.tail++
}

func (q *queue) not_empty() bool {
	return q.head < q.tail
}

func (q *queue) pop_front() int {
	q.head++
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
