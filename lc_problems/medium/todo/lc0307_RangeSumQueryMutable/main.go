type NumArray struct {
	nums []int
	sums []int
	n    int
}

func lowbit(x int) int {
	return x & (-x)
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	result := NumArray{
		nums: nums,
		sums: make([]int, n+1),
		n:    n,
	}
	for i := 0; i < n; i++ {
		result.update(i, nums[i])
	}
	return result
}

func (this *NumArray) update(index int, delta int) {
	index++
	n := this.n + 1
	for index < n {
		this.sums[index] += delta
		index += lowbit(index)
	}
}

func (this *NumArray) Update(index int, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	this.update(index, delta)
}

func (this *NumArray) sum(index int) int {
	index++
	result := 0
	for index > 0 {
		result += this.sums[index]
		index -= lowbit(index)
	}
	return result
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right) - this.sum(left-1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */