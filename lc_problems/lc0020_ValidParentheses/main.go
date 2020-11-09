package main

import "fmt"

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	ns := len(s)
	if ns%2 != 0 {
		return false
	}

	stack := make([]byte, 0, ns)

	for _, charRune := range s {
		char := byte(charRune)
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			ls := len(stack)
			if ls == 0 || stack[ls-1] != m[char] {
				return false
			}
			stack = stack[:ls-1]
		}
	}
	return len(stack) == 0
}

func main() {
	examples := "()"

	result := isValid(examples)

	fmt.Println(result)
}
