func isPossible(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, last := nums[0], nums[n-1]

	start := myAbs(first)
	arr := make([]int, start+last+2)

	for _, num := range nums {
		arr[start+num]++
	}

	if last-first < 2 {
		return false
	}
	fmt.Println(arr)

	for first <= last {
		count := 0
		for i := first; i <= last; i++ {
			if arr[i+start] == 0 {
				if count < 3 {
					return false
				} else {
					break
				}
			}
			if count >= 3 && arr[i-1+start] >= arr[i+start] {
				break
			}
			arr[i+start]--
			count++
		}

		if count < 3 {
			return false
		}

		for arr[first+start] == 0 && first <= last {
			first++
		}
	}
	return true
}

func myAbs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}