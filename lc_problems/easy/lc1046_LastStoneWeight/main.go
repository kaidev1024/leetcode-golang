import "container/heap"

type pq []int

func (this pq) Len() int {
	return len(this)
}

func (this pq) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this pq) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *pq) Push(x interface{}) {
	item := x.(int)
	*this = append(*this, item)
}

func (this *pq) Pop() interface{} {
	old := *this
	n := len(old)
	item := old[n-1]
	*this = old[0 : n-1]
	return item
}

func lastStoneWeight(stones []int) int {

	queue := pq(stones)
	heap.Init(&queue)

	for queue.Len() >= 2 {
		first := heap.Pop(&queue).(int)
		second := heap.Pop(&queue).(int)
		dif := first - second
		if dif > 0 {
			heap.Push(&queue, dif)
		}
	}
	if queue.Len() == 0 {
		return 0
	}
	return heap.Pop(&queue).(int)
}