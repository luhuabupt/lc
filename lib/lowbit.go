package lib

// 返回二进制最低位的1
func lowBit(x int) int {
	return x & -x
}
