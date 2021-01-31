func reorderSpaces(text string) string {
	nText := len(text)
	nSpaces := 0
	nWords := 0

	if text[0] == ' ' {
		nSpaces = 1
	}

	for i := 1; i < nText; i++ {
		if text[i] == ' ' {
			nSpaces++
			if text[i-1] != ' ' {
				nWords++
			}
		}
	}
	if text[nText-1] != ' ' {
		nWords++
	}

	separator := nSpaces
	if nWords != 1 {
		separator = nSpaces / (nWords - 1)
	}

	fmt.Println(nWords, separator, nSpaces, nText)
	result := make([]byte, nText)

	pos := 0
	rpos := 0
	countWords := 0

	for {
		for pos < nText && text[pos] == ' ' {
			pos++
		}

		if pos == nText {
			break
		}

		for pos < nText && text[pos] != ' ' {
			result[rpos] = text[pos]
			rpos++
			pos++
		}
		countWords++

		if countWords == nWords {
			break
		}
		for i := 0; i < separator; i++ {
			result[rpos] = ' '
			rpos++
		}
	}

	for rpos < nText {
		result[rpos] = ' '
		rpos++
	}

	return string(result)
}