package main

import (
	"fmt"
	"math/rand"
)

// generate-random-point-in-a-circle 在圆内随机生成点  2022-06-05 12:36:13
func main() {
	fmt.Println(generateRandomPointInACircle())
}

//leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	x, y, r float64
}

func Constructor(r float64, x float64, y float64) Solution {
	return Solution{x, y, r}
}

func (t *Solution) RandPoint() []float64 {
	for {
		a, b := rand.Float64()*2-1, rand.Float64()*2-1
		if a*a+b*b < 1 {
			return []float64{t.x + a*t.r, t.y + b*t.r}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
//leetcode submit region end(Prohibit modification and deletion)
