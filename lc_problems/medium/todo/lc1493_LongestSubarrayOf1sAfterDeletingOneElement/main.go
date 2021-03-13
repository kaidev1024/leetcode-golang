func longestSubarray(nums []int) int {
	globalZeros := 0
	zeros := 0
	result := 0
	prevOnes := 0
	curOnes := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			globalZeros++
			zeros++
			if zeros == 1 {
				ones := prevOnes + curOnes
				if ones > result {
					result = ones
				}
				prevOnes = curOnes
				curOnes = 0
			} else {
				prevOnes = 0
			}
		} else {
			zeros = 0
			curOnes++
		}
	}
	result = max(curOnes+prevOnes, result)
	if globalZeros == 0 {
		result--
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}