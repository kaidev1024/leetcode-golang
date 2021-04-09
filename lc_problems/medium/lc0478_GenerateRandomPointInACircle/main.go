\type Solution struct {
    r, x, y float64
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
    return Solution{radius, x_center, y_center}
}

func getNum() float64 {
    return (rand.Float64() - 0.5) * 2
}

func getPoint()(float64, float64) {
    x := getNum()
    y := getNum()
    if x * x + y * y <= 1 {
        return x, y
    }
    return getPoint()
}

func (this *Solution) RandPoint() []float64 {
    x, y := getPoint()
    x *= this.r
    y *= this.r
    return []float64{x + this.x, y + this.y}
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */