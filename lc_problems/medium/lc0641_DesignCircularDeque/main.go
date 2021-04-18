type MyCircularDeque struct {
	arr      []int
	capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	arr := make([]int, 0, k)
	return MyCircularDeque{
		arr:      arr,
		capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	arr := this.arr
	if len(arr) == this.capacity {
		return false
	}
	this.arr = append([]int{value}, arr...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	arr := this.arr
	if len(arr) == this.capacity {
		return false
	}
	this.arr = append(arr, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	arr := this.arr
	if len(arr) == 0 {
		return false
	}
	this.arr = arr[1:]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	arr := this.arr
	n := len(arr)
	if n == 0 {
		return false
	}
	this.arr = arr[:n-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	arr := this.arr
	if len(arr) == 0 {
		return -1
	}
	return arr[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	arr := this.arr
	if len(arr) == 0 {
		return -1
	}
	return arr[len(arr)-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.arr) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return len(this.arr) == this.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */