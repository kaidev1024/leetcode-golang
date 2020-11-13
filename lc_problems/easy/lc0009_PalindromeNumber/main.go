package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	xr := 0
	for x > xr {
		xr = xr*10 + x%10
		x /= 10
	}
	return xr == x || xr/10 == x
}

func main() {
	x1 := 100
	x2 := 121
	fmt.Println(isPalindrome(x1), isPalindrome(x2))
}
