package leetcode

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	left := 0
	right := l - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{0, 0}
}
