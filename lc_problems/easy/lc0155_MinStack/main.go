package leetcode

type MinStack struct {
	stack, minStack []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	l := len(this.minStack)
	if l == 0 || this.minStack[l-1] >= x {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	last := this.stack[l-1]
	this.stack = this.stack[:l-1]
	lmin := len(this.minStack)
	if this.minStack[lmin-1] == last {
		this.minStack = this.minStack[:lmin-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
