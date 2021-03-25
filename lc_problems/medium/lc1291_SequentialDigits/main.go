func getLow(low int) (int, int) {
	i := 10
	length := 2
	for {
		if low/i < 10 {
			break
		}
		i *= 10
		length++
	}
	curDigit := low / i
	if curDigit+length-1 > 9 {
		length++
		if length == 10 {
			return -1, -1
		}
		result := 0
		curDigit := 1
		for i := 0; i < length; i++ {
			result *= 10
			result += curDigit
			curDigit++
		}
		return result, length
	}
	result := 0
	for i := 0; i < length; i++ {
		result *= 10
		result += curDigit
		curDigit++
	}
	return result, length
}

func getIncrement(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return (result - 1) / 9
}

func getFirst(n int) int {
	result := 0
	curDigit := 1
	for i := 0; i < n; i++ {
		result *= 10
		result += curDigit
		curDigit++
	}
	return result
}

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0)
	curNumber, width := getLow(low)
	if width == -1 {
		return nil
	}
	increment := getIncrement(width)
	if curNumber < low {
		curNumber += increment
		if curNumber%10 == 0 {
			width++
			curNumber = getFirst(width)
			increment = getIncrement(width)
		}
	}
	for curNumber <= high {
		result = append(result, curNumber)
		curNumber += increment
		if curNumber%10 == 0 {
			width++
			curNumber = getFirst(width)
			increment = getIncrement(width)
		}
	}
	return result
}