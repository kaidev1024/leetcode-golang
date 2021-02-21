type FrontMiddleBackQueue struct {
	arr []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{make([]int, 0)}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.arr = append([]int{val}, this.arr...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	n := len(this.arr)
	half := append([]int{val}, this.arr[n/2:]...)
	this.arr = append(this.arr[:n/2], half...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.arr = append(this.arr, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	n := len(this.arr)
	if n == 0 {
		return -1
	}
	r := this.arr[0]
	this.arr = this.arr[1:]
	return r
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	n := len(this.arr)
	if n == 0 {
		return -1
	}
	r := this.arr[(n-1)/2]
	this.arr = append(this.arr[:(n-1)/2], this.arr[(n-1)/2+1:]...)
	return r
}

func (this *FrontMiddleBackQueue) PopBack() int {
	n := len(this.arr)
	if n == 0 {
		return -1
	}
	r := this.arr[n-1]
	this.arr = this.arr[:n-1]
	return r
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */