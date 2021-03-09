//O(nlgn)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, 1, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		index := sort.SearchInts(dp, nums[i])
		if index == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[index] = nums[i]
		}
	}
	return len(dp)
}

//O(n^2)

// func lengthOfLIS(nums []int) int {
//     n := len(nums)
//     dp := make([]int, n)
//     dp[0] = 1
//     result := 1

//     for i := 1; i < n; i++ {
//         maxVal := 0
//         for j := 0; j < i; j++ {
//             if nums[i] > nums[j] {
//                 maxVal = max(maxVal, dp[j])
//             }
//         }
//         dp[i] = maxVal + 1
//         result = max(dp[i], result)
//     }
//     return result
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }