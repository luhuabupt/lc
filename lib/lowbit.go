package main

import "fmt"

func main() {
	gospersHack(4, 0)
	gospersHack(4, 1)
	gospersHack(4, 2)
	gospersHack(4, 3)
	gospersHack(4, 4)
}

// 返回二进制最低位的1
func lowBit(x int) int {
	return x & -x
}

// 按二进制1的位数枚举
// @see https://zhuanlan.zhihu.com/p/360512296
func gospersHack(num, kth int) {
	mask := (1 << kth) - 1
	limit := 1 << num
	for mask < limit {
		fmt.Printf("%b \n", mask)

		if mask == 0 {
			break
		}
		lowBit := mask & -mask
		right := mask + lowBit
		mask = (((right ^ mask) >> 2) / lowBit) | right
	}
}
