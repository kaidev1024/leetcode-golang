import "strconv"

func diffWaysToCompute(input string) []int {

	opChars := map[byte]func(int, int) int{
		'+': func(a, b int) int { return a + b },
		'-': func(a, b int) int { return a - b },
		'*': func(a, b int) int { return a * b },
	}

	m := make(map[string][]int)

	var helper func(input string) []int
	helper = func(input string) []int {
		res := []int{}
		if arr, ok := m[input]; ok {
			return arr
		}

		for i := range input {
			if op, ok := opChars[input[i]]; ok {
				leftRes := helper(input[:i])
				rightRes := helper(input[i+1:])

				for _, leftV := range leftRes {
					for _, rightV := range rightRes {
						res = append(res, op(leftV, rightV))
					}
				}
			}
		}
		if len(res) == 0 {
			num, _ := strconv.Atoi(input)
			res = []int{num}
		}
		m[input] = res
		return res
	}
	helper(input)
	return m[input]
}