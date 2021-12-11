package main

import "sort"

type Robot struct {
	w, h int
	x, y int
	dir  string
}

func Constructor(width int, height int) Robot {
	return Robot{
		w:   width,
		h:   height,
		x:   0,
		y:   0,
		dir: "East",
	}
}

func (this *Robot) Move(num int) {
	circle := 2*this.w + 2*this.h - 4
	num %= circle
	if num == 0 {
		return
	}
	//pos := map[string][2]int{"East": {1, 0}, "North": {0, 1}, "West": {-1, 0}, "South": {0, -1},}
	if this.dir == "East" {
		if this.x+num < this.w-1 {
			this.x += num
		} else {
			this.x = this.w - 1
			this.dir = "North"
			this.Move(num - (this.w - 1 - this.x))
		}
	} else if this.dir == "North" {
		if this.y+num < this.h-1 {
			this.y += num
		} else {
			this.y = this.h - 1
			this.dir = "West"
			this.Move(num - (this.h - 1 - this.y))
		}
	} else if this.dir == "West" {
		if this.x-num > 0 {
			this.x -= num
		} else {
			this.x = 0
			this.dir = "South"
			this.Move(num - (this.x - 0))
		}
	} else if this.dir == "South" {
		if this.y-num > 0 {
			this.y -= num
		} else {
			this.y = 0
			this.dir = "East"
			this.Move(num - (this.y - 0))
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	return this.dir
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0] || items[i][0] == items[j][0] && items[i][1] > items[j][1]
	})

	a, b := []int{}, []int{}
	max := 0
	for _, arr := range items {
		if arr[1] > max {
			max = arr[1]
		}
		a = append(a, arr[0])
		b = append(b, max)
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		//x := sort.Search(int(1E9), func(v int) bool {
		//	return  stack[v][0] >= queries[i]
		//})
		//ans[i] = stack[x][1]
		x := sort.SearchInts(a, queries[i])
		ans[i] = b[x]
	}

	return ans
}
