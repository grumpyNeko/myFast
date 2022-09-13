package myFast

func arrayToBST(data []uint32) []uint32 {
	size := len(data)
	if !isFullBSTSize(size) {
		panic("given size cannot apply to a full bst")
	}
	tree := make([]uint32, size, size)

	setNode(0, data, tree)
	return tree
}

// size should be `1<<n - 1`
func isFullBSTSize(size int) bool {
	for i := 0; i < 33; i++ {
		if size == 1<<i-1 {
			return true
		}
	}
	return false
}

/*
tree[nodePos] = data[mid]
set left child
set right child
*/
func setNode(nodePos int, data []uint32, tree []uint32) {
	size := len(data)
	if size == 0 {
		return
	}
	if size == 1 {
		tree[nodePos] = data[0]
	}
	if size == 2 { // should not be 2n actually
		panic("should not be 2!")
	}
	mid := size >> 1
	tree[nodePos] = data[mid]
	l := left(nodePos)
	if l < len(tree) { // todo: i wonder whether this check is necessary
		setNode(l, data[:mid], tree)
	}
	r := right(nodePos)
	if r < len(tree) {
		setNode(r, data[mid+1:], tree) // todo: mid+1 ???
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func getFromBst(bst []uint32, target uint32) bool {
	cur := 0
retry:
	flag := target - bst[cur]
	if flag == 0 {
		return true
	} else {
		if right(cur) >= len(bst) { // if don't have right, also don't have left
			return false
		}
		if flag > 0 {
			cur = right(cur)
		} else {
			cur = left(cur)
		}
		goto retry
	}
}
