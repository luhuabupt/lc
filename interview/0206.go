package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(sortByAge([]user{{"a", 10}, {"b", 20}, {"c", 8}, {"d", 10}}))
	//fmt.Println(sortByAge([]user{}))
	//fmt.Println(sortByAge([]user{{"aa", 10000}}))
	printListNode(reverseListNode(formatListNode([]int{1, 2, 3})))
	printListNode(formatListNode([]int{1, 2, 3}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println("")
}

func formatListNode(a []int) *ListNode {
	p := &ListNode{}
	head := &ListNode{0, nil}
	head.Next = p
	for i, v := range a {
		p.Val = v
		if i < len(a)-1 {
			p.Next = &ListNode{0, nil}
		}
		p = p.Next
	}

	return head.Next
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next                 // 反转前下一个节点
	ans := reverseListNode(head.Next) // 反转后的头节点

	// 将当前节点加到返回后的尾节点上
	head.Next = nil
	next.Next = head

	return ans
}

type user struct {
	name string
	age  int
}

func sortByAge(users []user) []user {
	sort.Slice(users, func(i, j int) bool {
		return users[i].age < users[j].age || users[i].age == users[j].age && users[i].name < users[j].name
	})

	return users
}
