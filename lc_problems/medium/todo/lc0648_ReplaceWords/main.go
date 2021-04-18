func replaceWords(dict []string, sentence string) string {
	trie := Constructor()
	for _, v := range dict {
		trie.Insert(v)
	}
	strs := split(sentence)
	for i := 0; i < len(strs); i++ {
		if s := trie.Search(strs[i]); s != "" {
			strs[i] = s
		}
	}
	return join(strs)
}

type Trie struct {
	m map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		m: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for i := 0; i < len(word); i++ {
		w := word[i]
		if v, ok := current.m[w]; ok {
			current = v
		} else {
			current.m[w] = &Trie{m: make(map[byte]*Trie)}
			current = current.m[w]
		}
	}
	// represents the end of the string
	current.m['0'] = &Trie{m: make(map[byte]*Trie)}
}

func (this *Trie) Search(word string) string {
	current, res := this, ""
	for i := 0; i < len(word); i++ {
		v, ok := current.m[word[i]]
		if !ok {
			return ""
		}

		res = res + string(word[i])
		// return result if the end of the string
		if v.m['0'] != nil {
			return res
		}
		current = v
	}
	return res
}

func split(str string) []string {
	res, i, s := []string{}, 0, 0
	for i < len(str) {
		if str[i] == ' ' {
			res = append(res, str[s:i])
			s = i + 1
		}
		i++
	}
	return append(res, str[s:i])
}

func join(s []string) string {
	res := ""
	for _, v := range s {
		if res == "" {
			res = v
		} else {
			res = res + " " + v
		}
	}
	return res
}