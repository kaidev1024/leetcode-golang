

func calculate(s string) int {
	n := len(s)
	pos := 0

	skipSpace := func() {
		for pos < n {
			if s[pos] == ' ' {
				pos++
			} else {
				break
			}
		}
	}

	parseInt := func() int {
		r := 0
		for pos < n {
			c := s[pos]
			if c >= '0' && c <= '9' {
				r = r*10 + int(c-'0')
				pos++
			} else {
				break
			}
		}
		return r
	}

	nums := make([]int, 0)
	ops := make([]byte, 0)
	operate := func(op byte) {
		n := len(nums)
		a := nums[n-2]
		b := nums[n-1]
		nums = nums[:n-2]
		if op == '*' {
			nums = append(nums, a*b)
			return
		}
		nums = append(nums, a/b)
	}

	skipSpace()
	if pos == n {
		return 0
	}
	nums = append(nums, parseInt())

	for pos < n {

		skipSpace()
		if pos == n {
			break
		}
		op := s[pos]
		pos++

		skipSpace()
		if pos == n {
			break
		}
		nums = append(nums, parseInt())

		if op == '*' || op == '/' {
			operate(op)
		} else {
			ops = append(ops, op)
		}
	}

	nops := len(ops)
	result := nums[0]
	for i := 0; i < nops; i++ {
		v := nums[i+1]
		op := ops[i]
		if op == '+' {
			result += v
		} else {
			result -= v
		}
	}
	return result
}