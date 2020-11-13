package main

import "fmt"

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)

	if la > lb {
		return addBinary(b, a)
	}

	var carry byte

	result := make([]byte, lb)

	for i := lb - 1; i >= 0; i-- {
		ia := i - lb + la

		sum := (b[i] - '0') + byte(carry)
		if ia >= 0 {
			sum += a[ia] - '0'
		}
		carry = sum >> 1
		sum &= 1
		result[i] = '0' + sum
	}

	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}
	return string(result)
}

func main() {
	a := "11"
	b := "1"
	result := addBinary(a, b)
	fmt.Println(result)
}
