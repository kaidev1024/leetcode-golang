type OrderedStream struct {
	pos    int
	length int
	data   []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{1, n + 1, make([]string, n+1)}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.data[id] = value
	if this.data[this.pos] == "" {
		return nil
	}
	result := make([]string, 0, this.length)
	for this.pos < this.length && this.data[this.pos] != "" {
		result = append(result, this.data[this.pos])
		this.pos++
	}
	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */