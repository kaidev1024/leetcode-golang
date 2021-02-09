func isLongPressedName(name string, typed string) bool {
	nn := len(name)
	nt := len(typed)
	if nt < nn {
		return false
	}

	pos := 0
	for i := 0; i < nn; {
		if pos == nt {
			return false
		}
		if name[i] == typed[pos] {
			i++
			pos++
			continue
		}
		if i == 0 {
			return false
		}
		if typed[pos] != typed[pos-1] {
			return false
		}
		pos++
		if pos == nt {
			return false
		}
	}
	for pos < nt {
		if typed[pos] != typed[pos-1] {
			return false
		}
		pos++
	}
	return true
}