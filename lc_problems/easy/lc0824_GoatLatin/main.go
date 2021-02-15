func toGoatLatin(S string) string {
	isVowel := func(character uint8) bool {
		return character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' ||
			character == 'A' || character == 'E' || character == 'I' || character == 'O' || character == 'U'
	}

	var a, b = []uint8{}, []uint8{}

	for i, j := 0, 0; j < len(S); j++ {
		if S[j] == ' ' || j == len(S)-1 {

			isEnd := j == len(S)-1

			if isEnd {
				j++
			}
			switch {
			case isVowel(S[i]):
				{
					b = append(b, S[i:j]...)

					break
				}
			case !isVowel(S[i]):
				{
					b = append(b, S[i+1:j]...)
					b = append(b, S[i])
					break
				}
			}
			b = append(b, "ma"[:]...)
			a = append(a, 'a')
			b = append(b, a...)
			if !isEnd {
				b = append(b, ' ')
			}

			i = j + 1
		}
	}
	return string(b)
}