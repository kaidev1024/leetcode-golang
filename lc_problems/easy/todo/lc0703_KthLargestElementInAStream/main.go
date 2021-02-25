type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	r := old[n-1]
	*h = old[:n-1]
	return r
}

type KthLargest struct {
	h *MaxHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	mh := MaxHeap(nums)
	fmt.Println(mh)
	h := make(MaxHeap, 0)
	return KthLargest{&h, k}
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(this.h, val)
	}
	top := heap.Pop(this.h).(int)
	if top < val {
		top = val
	}
	heap.Push(this.h, top)
	return top
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */