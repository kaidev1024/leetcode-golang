package leetcode

func majorityElement(nums []int) int {
	n := len(nums)
	currentElement := nums[0]
	count := 1

	for i := 1; i < n; i++ {
		if count == 0 {
			count++
			currentElement = nums[i]
		} else if currentElement == nums[i] {
			count++
		} else {
			count--
		}
	}
	return currentElement
}
