func majorityElement(nums []int) []int {

	var result []int
	num1 := math.MaxInt16
	num2 := math.MaxInt16

	count1 := 0
	count2 := 0

	for _, v := range nums {
		if num1 == v {
			count1++
		} else if num2 == v {
			count2++
		} else if count1 == 0 {
			num1 = v
			count1++
		} else if count2 == 0 {
			num2 = v
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0

	for _, v := range nums {
		if v == num1 {
			count1++
		} else if v == num2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, num1)
	}
	if count2 > len(nums)/3 {
		result = append(result, num2)
	}
	return result
}
