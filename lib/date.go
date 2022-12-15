package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(y2i("2022-01-02 15:04:05"))
}

// "01:02"
func m2i(v string) int {
	arr := strings.Split(v, ":")
	h, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	return h*60 + m
}

// "01:02:03"
func s2i(v string) int {
	arr := strings.Split(v, ":")
	h, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	s, _ := strconv.Atoi(arr[2])
	return h*60*60 + m*60 + s
}

// "1997-10-10 23:59:59"
func y2i(v string) int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", v)
	return t.Unix()
}
