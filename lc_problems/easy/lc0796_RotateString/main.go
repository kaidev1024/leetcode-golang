import "strings"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	A = A + A
	return strings.Contains(A, B)
}