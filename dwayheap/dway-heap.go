package dwayheap

type Item struct {
	Elem     any
	Priority int
}

// DHeap to implemented as max heap having the biggest number at the top
type DHeap struct {
	d int // d-way

	data []Item
}

func (d *DHeap) parentIndex(index int) int {
	return (index - 1) / d.d
}

func (d *DHeap) Top() {

}

func (d *DHeap) Peek() {

}

func (d *DHeap) Insert(item Item) {

}

func (d *DHeap) Remove(e any) {

}

func (d *DHeap) Update() {

}

func (d *DHeap) bubbleUp(index int) {
	parentIdx := index
	for parentIdx > 0 {
		currIdx := parentIdx
		parentIdx = d.parentIndex(index)
		// find the proper position. at worst case, the position will be zero.
		if d.data[parentIdx].Priority > d.data[currIdx].Priority {
			break
		}
		d.swap(currIdx, parentIdx)
	}
}

func (d *DHeap) pushDown() {

}

func (d *DHeap) swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}
