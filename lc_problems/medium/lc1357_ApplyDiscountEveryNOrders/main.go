type Cashier struct {
	bills    int
	n        int
	discount int
	m        map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	m := make(map[int]int)
	for i, p := range products {
		m[p] = prices[i]
	}
	return Cashier{0, n, discount, m}
}

func (cashier *Cashier) GetBill(product []int, amount []int) float64 {
	cashier.bills++

	var sum float64 = 0
	for i, p := range product {
		sum += float64(cashier.m[p] * amount[i])
	}

	if cashier.bills == cashier.n {
		cashier.bills = 0
		return sum - (float64(cashier.discount)*sum)/100
	}

	return sum
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */