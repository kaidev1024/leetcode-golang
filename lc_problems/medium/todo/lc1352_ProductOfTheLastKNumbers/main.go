type ProductOfNumbers struct {
	products []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		products: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.products = []int{1}
	} else {
		this.products = append(this.products, num*this.products[len(this.products)-1])
	}

}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.products)
	if n-1 < k {
		return 0
	}
	return this.products[n-1] / this.products[n-k-1]
}

