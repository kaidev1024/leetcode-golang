import "strings"

func split(s string) (string, string) {
	pos := strings.Index(s, " ")
	head := s[:pos-1]
	body := s[pos+1:]
	return head, body
}

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		headi, bodyi := split(logs[i])
		headj, bodyj := split(logs[j])
		isDigiti := false
		if bodyi[0] < 65 {
			isDigiti = true
		}
		isDigitj := false
		if bodyj[0] < 65 {
			isDigitj = true
		}

		if isDigiti && isDigitj {
			return i < j
		}
		if isDigiti && !isDigitj {
			return false
		}

		if !isDigiti && isDigitj {
			return true
		}

		if bodyi != bodyj {
			return bodyi < bodyj
		}
		return headi < headj
	})

	return logs
}