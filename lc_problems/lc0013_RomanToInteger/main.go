package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	ls := len(s)
	result := m[s[ls-1]]
	latest := result
	for i := ls - 2; i >= 0; i-- {
		val := m[s[i]]
		if val < latest {
			result -= val
			continue
		}
		result += val
		latest = val
	}
	return result
}

func main() {
	s1 := "III"
	s2 := "IV"
	fmt.Println(romanToInt(s1), romanToInt(s2))
}
