package main

import "fmt"

func lengthOfLastWord(s string) int {
	l := len(s)
	result := 0
	i := l - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	if i < 0 {
		return 0
	}

	for ; i >= 0; i-- {
		if s[i] == ' ' {
			return result
		}
		result++
	}
	return result
}

func main() {
	word := "Hello World"
	result := lengthOfLastWord(word)
	fmt.Println(result)
}
