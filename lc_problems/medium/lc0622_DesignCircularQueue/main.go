type MyCircularQueue struct {
	cap int
	arr []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap: k,
		arr: make([]int, 0, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	arr := this.arr
	if len(arr) == this.cap {
		return false
	}
	this.arr = append(arr, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	arr := this.arr
	if len(arr) == 0 {
		return false
	}
	this.arr = arr[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	n := len(this.arr)
	if n == 0 {
		return -1
	}
	return this.arr[0]
}

func (this *MyCircularQueue) Rear() int {
	n := len(this.arr)
	if n == 0 {
		return -1
	}
	return this.arr[n-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.arr) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.arr) == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */