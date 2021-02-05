

func numPrimeArrangements(n int) int {
	m := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	var primeNumbersCount uint64
	for i := 0; i < len(m) && m[i] <= n; i++ {
		primeNumbersCount++
	}

	return int(FactorialMemoization(primeNumbersCount)*FactorialMemoization(uint64(n)-primeNumbersCount)) % 1000000007

}

const LIM = 100

var facts [LIM]uint64

func FactorialMemoization(n uint64) (res uint64) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = (n * FactorialMemoization(n-1)) % 1000000007
		return res
	}

	return 1
}