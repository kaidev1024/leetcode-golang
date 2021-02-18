func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	var overlap int
	if A >= G || E >= C || B >= H || F >= D {
		overlap = 0
	} else {
		overlap = abs(max(A, E), min(C, G)) * abs(max(B, F), min(D, H))
	}

	area1 := abs(A, C) * abs(B, D)
	area2 := abs(E, G) * abs(F, H)

	return area1 + area2 - overlap
}