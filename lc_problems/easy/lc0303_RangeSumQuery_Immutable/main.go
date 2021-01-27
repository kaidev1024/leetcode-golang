package leetcode

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	sums := make([]int, l+1)
	for i := 1; i <= l; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	return NumArray{sums}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sums[j+1] - this.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
