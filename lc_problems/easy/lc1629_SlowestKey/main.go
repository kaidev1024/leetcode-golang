func slowestKey(releaseTimes []int, keysPressed string) byte {
	result := releaseTimes[0]
	b := keysPressed[0]
	l := len(releaseTimes)
	for i := 1; i < l; i++ {
		d := releaseTimes[i] - releaseTimes[i-1]
		if d > result {
			result = d
			b = keysPressed[i]
		}
		if d == result {
			if b < keysPressed[i] {
				b = keysPressed[i]
			}
		}
	}
	return b
}