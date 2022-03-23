package main

import "fmt"

func main() {
	obj := Constructor()
	//obj.Inc("hello")
	//fmt.Println("get: ", obj.GetMaxKey(), "|", obj.GetMinKey())
	//obj.Inc("hello")
	//fmt.Println("get: ", obj.GetMaxKey(), "|", obj.GetMinKey())
	//obj.Inc("leet")
	//fmt.Println("get: ", obj.GetMaxKey(), "|", obj.GetMinKey())
	//obj.Dec("hello")
	//fmt.Println("get: ", obj.GetMaxKey(), "|", obj.GetMinKey())
	//obj.Dec("hello")
	//fmt.Println("get: ", obj.GetMaxKey(), "|", obj.GetMinKey())
	//obj.Dec("leet")
	//obj.Inc("leet")
	//fmt.Println("get: ", obj.GetMaxKey(), "|", obj.GetMinKey())
	a := []string{"a", "b", "b", "c", "c", "c", "b", "b", "a"}
	b := []int{1, 1, 1, 1, 1, 1, 0, 0, 0}

	for i, x := range b {
		if x == 1 {
			obj.Inc(a[i])
		} else {
			obj.Dec(a[i])
		}
		fmt.Println(obj.GetMaxKey(), "|", obj.GetMinKey())
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
type AllOne struct {
	h   map[string]*Node
	dll *doubleLinkList
}

type doubleLinkList struct {
	head *Node
	tail *Node
}

type Node struct {
	Next *Node
	Pre  *Node
	Key  map[string]bool
	Val  int
}

func Constructor() AllOne {
	return AllOne{
		h: map[string]*Node{},
		dll: &doubleLinkList{
			head: nil,
			tail: nil,
		},
	}
}

func (t *AllOne) Inc(key string) {
	// init
	if t.h[key] == nil {
		if t.dll.head == nil {
			ne := &Node{
				nil,
				nil,
				map[string]bool{key: true},
				1,
			}
			t.dll.head = ne
			t.dll.tail = ne
			t.h[key] = ne
		} else if t.dll.head.Val == 1 {
			t.dll.head.Key[key] = true
			t.h[key] = t.dll.head
		} else {
			ne := &Node{
				t.dll.head,
				nil,
				map[string]bool{key: true},
				1,
			}
			t.dll.head.Pre = ne
			t.dll.head = ne
			t.h[key] = ne
		}
		return
	}

	delete(t.h[key].Key, key)
	pre := t.h[key]

	if pre.Next == nil {
		pre.Next = &Node{
			nil,
			pre,
			map[string]bool{key: true},
			pre.Val + 1,
		}
		t.dll.tail = pre.Next
	} else if pre.Next.Val > pre.Val+1 {
		ne := &Node{
			pre.Next,
			pre,
			map[string]bool{key: true},
			pre.Val + 1,
		}
		pre.Next.Pre = ne
		pre.Next = ne
	} else {
		pre.Next.Key[key] = true
	}
	t.h[key] = pre.Next

	if len(pre.Key) == 0 {
		if t.dll.head == pre {
			t.dll.head = pre.Next
		}
		delNode(pre)
	}
}

func (t *AllOne) Dec(key string) {
	cur := t.h[key]
	delete(cur.Key, key)

	//fmt.Println("dec", key)
	if cur.Val == 1 {
		t.h[key] = nil
	} else if cur.Pre == nil {
		//fmt.Println("pre nil ", key)
		ne := &Node{
			cur,
			nil,
			map[string]bool{key: true},
			cur.Val - 1,
		}
		cur.Pre = ne
		t.dll.head = ne
		t.h[key] = ne
	} else if cur.Pre.Val == cur.Val-1 {
		//fmt.Println("cpv", key)
		cur.Pre.Key[key] = true
		t.h[key] = cur.Pre
	} else {
		//fmt.Println("else ")
		ne := &Node{
			cur,
			cur.Pre,
			map[string]bool{key: true},
			cur.Val - 1,
		}
		cur.Pre.Next = ne
		cur.Pre = ne
		t.h[key] = ne
	}

	if len(cur.Key) == 0 {
		if t.dll.tail == cur {
			t.dll.tail = cur.Pre
		}
		if t.dll.head == cur {
			t.dll.head = t.dll.head.Next
		}

		delNode(cur)
	}

	//fmt.Println(t.dll, t.h[key])
}

func (t *AllOne) GetMaxKey() string {
	if t.dll.tail == nil {
		return ""
	}
	for x, _ := range t.dll.tail.Key {
		return x
	}
	return ""
}

func (t *AllOne) GetMinKey() string {
	if t.dll.head == nil {
		return ""
	}
	for x, _ := range t.dll.head.Key {
		return x
	}
	return ""
}

func delNode(p *Node) {
	if p.Pre != nil {
		p.Pre.Next = p.Next
	}
	if p.Next != nil {
		p.Next.Pre = p.Pre
	}
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
//leetcode submit region end(Prohibit modification and deletion)
