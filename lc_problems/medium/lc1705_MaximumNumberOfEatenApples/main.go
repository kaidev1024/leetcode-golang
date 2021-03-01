import "container/heap"

type MyHeap [][]int

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	r := old[n-1]
	*h = old[:n-1]
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func eatenApples(apples []int, days []int) int {
	mh := &MyHeap{}
	heap.Init(mh)

	result := 0
	n := len(apples)
	for i := 0; i < n; i++ {
		if apples[i] > 0 {
			heap.Push(mh, []int{apples[i], days[i] + i})
		}
		if mh.Len() == 0 {
			continue
		}
		item := heap.Pop(mh).([]int)
		for mh.Len() > 0 && item[1] <= i {
			item = heap.Pop(mh).([]int)
		}

		if item[1] >= 1+i {
			result++
			item[0]--
			if item[0] > 0 && item[1] > i+1 {
				heap.Push(mh, item)
			}
		}

	}

	day := n
	for mh.Len() > 0 {
		item := heap.Pop(mh).([]int)
		napple := min(item[0], item[1]-day)
		result += napple
		day += napple
	}

	return result
}