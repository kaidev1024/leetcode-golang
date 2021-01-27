package leetcode

func moveZeroes(nums []int) {
	l := len(nums)
	slow := 0
	fast := 0
	for fast < l {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
			fast++
		} else {
			fast++
		}
	}
	for slow < l {
		nums[slow] = 0
		slow++
	}
}
