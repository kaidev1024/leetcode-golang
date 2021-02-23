func sortColors(nums []int) {
	var left, right int

	for left = 0; left < len(nums) && nums[left] == 0; left++ {
	}
	for right = len(nums) - 1; right >= 0 && nums[right] == 2; right-- {
	}

	curr := left
	for curr <= right {
		switch nums[curr] {
		case 0:
			nums[curr] = nums[left]
			nums[left] = 0
			left++
			curr++
		case 1:
			curr++
		case 2:
			nums[curr] = nums[right]
			nums[right] = 2
			right--
		}
	}
}