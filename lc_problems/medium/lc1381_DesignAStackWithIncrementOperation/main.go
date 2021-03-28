type CustomStack struct {
	maxSize int
	arr     []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		maxSize: maxSize,
		arr:     make([]int, 0, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.arr) == this.maxSize {
		return
	}
	this.arr = append(this.arr, x)
}

func (this *CustomStack) Pop() int {
	n := len(this.arr)
	if n == 0 {
		return -1
	}
	ret := this.arr[n-1]
	this.arr = this.arr[:n-1]
	return ret
}

func (this *CustomStack) Increment(k int, val int) {
	n := len(this.arr)
	if n < k {
		k = n
	}
	for i := 0; i < k; i++ {
		this.arr[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */