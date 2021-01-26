package leetcode

func isIsomorphic(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}
	mapping := make(map[byte]byte)
	set := make(map[byte]struct{})

	for i := 0; i < l; i++ {
		si := s[i]
		ti := t[i]
		if sm, ok := mapping[si]; ok {
			if sm != ti {
				return false
			}
			continue
		}
		if _, ok := set[ti]; ok {
			return false
		}

		mapping[si] = ti
		set[ti] = struct{}{}
	}
	return true
}
