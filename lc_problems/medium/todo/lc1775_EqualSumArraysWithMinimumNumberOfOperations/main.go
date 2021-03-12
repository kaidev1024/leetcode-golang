func helper(nums []int) (n, sum int, counts []int) {
	n = len(nums)
	counts = make([]int, 7)
	for _, v := range nums {
		sum += v
		counts[v]++
	}
	return
}

func minOperations(nums1 []int, nums2 []int) int {
	n1, sum1, counts1 := helper(nums1)
	n2, sum2, counts2 := helper(nums2)
	if sum1 == sum2 {
		return 0
	}
	if n1 > n2*6 || n2 > n1*6 {
		return -1
	}

	if sum1 > sum2 {
		sum1, sum2 = sum2, sum1
		counts1, counts2 = counts2, counts1
	}

	ret := 0

	for i := 1; i <= 6; i++ {
		ci := counts1[i]
		dif := 6 - i
		difs := dif * ci
		if sum1+difs >= sum2 {
			return ret + (sum2-sum1-1)/dif + 1
		} else {
			sum1 += difs
			ret += ci
		}

		ci = counts2[7-i]
		difs = dif * ci
		if sum2-difs <= sum1 {
			return ret + (sum2-sum1-1)/dif + 1
		} else {
			sum2 -= difs
			ret += ci
		}
	}
	return -1
}