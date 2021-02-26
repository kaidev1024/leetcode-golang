func minPartitions(n string) int {
	var r byte = '0'
	for i, l := 0, len(n); i < l; i++ {
		if r < n[i] {
			r = n[i]
		}
	}
	return int(r - '0')
}