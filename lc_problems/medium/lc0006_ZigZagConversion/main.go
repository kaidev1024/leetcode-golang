import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := make([]strings.Builder, numRows)
	pos := 0
	n := len(s)

	for pos < n {
		i := 0
		for ; i < numRows-1 && pos < n; i++ {
			arr[i].WriteByte(s[pos])
			pos++
		}
		for ; i > 0 && pos < n; i-- {
			arr[i].WriteByte(s[pos])
			pos++
		}
	}

	result := ""
	for i := 0; i < numRows; i++ {
		result += arr[i].String()
	}
	return result
}