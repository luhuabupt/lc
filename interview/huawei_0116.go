package main

import "fmt"

//终端公司的零售店需要定期去仓库提取货物，假设零售店和仓库在一个矩阵上：相邻点的距离为 1 ；只能通过上下左右走动。
//矩阵元素的值仅为三种：0，表示仓库； -1，表示障碍； 1，表示零售店。 注：障碍无法通过，其它可以通过。
//为了将取货效率最大化，需要计算每个零售店走到最近仓库的最小距离，并输出这些最小距离的和：
//无法到达仓库的零售店，不参与距离和的计算；
//没有零售店或者没有仓库的话，返回0;
//解答要求
//时间限制：3000ms, 内存限制：256MB
//输入
//一个 m*n 的数组（ m和n范围为 [1,300) ）
//
//输出：一个整数，表示所计算的最小距离和。
//
//样例
//输入
//1 -1 0
//0 1 1
//1 -1 1
//输出
//6
//样例解释：
//共有5个零售店（数字1），2个仓库（数字0），2个障碍（数字-1），零售店上标注的数字表示到最近仓库的最小距离：
//位置[2][2]的零售店，距离[0][2]的仓库为 2，距离[1][0]的仓库为 3，因此到最近仓库的最小距离为 2 ； 其余零售店到最近仓库的最小距离都为 1。所以，所有零售店到仓库的最小距离和为 1 + 1 + 1 + 1 + 2 = 6
func main() {
	fmt.Println(cal([][]int{{1, -1, 0}, {0, 1, 1}, {1, -1, 1}}))
	fmt.Println(cal([][]int{{1, 1, 0}, {1, 1, 1}, {1, 1, 1}}))
}

func cal(g [][]int) int {
	m := len(g)
	n := len(g[0])

	vis := make([][]bool, m) // 是否遍历过
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	queue := [][]int{} // 已经找到仓库的商店队列
	for i, row := range g {
		for j, v := range row {
			if v == 1 {
				for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
					x, y := i+d[0], j+d[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if g[x][y] == 0 {
							vis[i][j] = true
							queue = append(queue, []int{i, j}) // 队列初始化
							break
						}
					}
				}
			}
		}
	}

	ans := len(queue)
	for deep := 2; len(queue) > 0; deep++ {
		nq := [][]int{}
		for _, v := range queue {
			for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				x, y := v[0]+d[0], v[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if vis[x][y] {
						continue
					}
					if g[x][y] == 1 { // 找到当前队列周围节点，但尚未访问的节点：写入队列，计算距离
						vis[x][y] = true
						nq = append(nq, []int{x, y})
						ans += deep
					}
				}
			}
		}
		queue = nq // 队列重置
	}

	return ans
}
