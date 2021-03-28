type Cashier struct {
	target   int
	count    int
	discount int
	prices   map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	m := make(map[int]int)
	for i, product := range products {
		m[product] = prices[i]
	}
	return Cashier{
		target:   n,
		count:    0,
		discount: discount,
		prices:   m,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.count++
	total := 0
	for i, p := range product {
		total += this.prices[p] * amount[i]
	}

	result := float64(total)
	if this.count == this.target {
		this.count = 0
		result = float64(total) * (1 - float64(this.discount)/100)
	}
	return result
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */