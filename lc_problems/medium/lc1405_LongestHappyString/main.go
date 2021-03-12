import "container/heap"

type entry struct {
	count int
	char  byte
}

type PriorityQueue []*entry

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(v interface{}) {
	*pq = append(*pq, v.(*entry))
}
func (pq *PriorityQueue) Pop() interface{} {
	oldArr := *pq
	n := len(oldArr)
	lastEntry := oldArr[n-1]
	*pq = oldArr[:n-1]
	return lastEntry
}

func longestDiverseString(a int, b int, c int) string {
	pq := &PriorityQueue{}
	if a > 0 {
		heap.Push(pq, &entry{a, 'a'})
	}
	if b > 0 {
		heap.Push(pq, &entry{b, 'b'})
	}
	if c > 0 {
		heap.Push(pq, &entry{c, 'c'})
	}

	heap.Init(pq)
	var sb strings.Builder
	var prevChar byte
	for pq.Len() > 0 {
		e1 := heap.Pop(pq).(*entry)
		if e1.count == 1 {
			sb.WriteByte(e1.char)
			continue
		}
		if e1.count == 2 {
			if e1.char == prevChar {
				sb.WriteByte(e1.char)
				e1.count = 1
			} else {
				sb.WriteByte(e1.char)
				sb.WriteByte(e1.char)
				prevChar = e1.char
				continue
			}
		} else {
			if e1.char == prevChar {
				sb.WriteByte(e1.char)
				e1.count -= 1
			} else {
				sb.WriteByte(e1.char)
				sb.WriteByte(e1.char)
				e1.count -= 2
			}
		}

		if pq.Len() == 0 {
			return sb.String()
		}

		e2 := heap.Pop(pq).(*entry)
		heap.Push(pq, e1)
		sb.WriteByte(e2.char)
		e2.count -= 1
		if e2.count > 0 {
			heap.Push(pq, e2)
		}
		prevChar = e2.char
	}
	return sb.String()
}