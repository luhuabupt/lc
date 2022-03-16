package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	fmt.Println(obj.GetMinKey(), obj.GetMaxKey())
}

//leetcode submit region begin(Prohibit modification and deletion)
// 双链表
type node struct {
	keys map[string]bool
	val  int
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{
		list.New(),
		map[string]*list.Element{},
	}
}

func (t *AllOne) Inc(key string) {
	cur := t.nodes[key]

	if cur == nil {
		if t.Front() == nil || t.Front().Value.(node).val > 1 {
			t.nodes[key] = t.PushFront(node{map[string]bool{key: true}, 1})
		} else {
			t.Front().Value.(node).keys[key] = true
			t.nodes[key] = t.Front()
		}
		return
	}

	nxt := cur.Next()
	cn := cur.Value.(node)
	if nxt == nil || nxt.Value.(node).val > cn.val+1 {
		t.nodes[key] = t.InsertAfter(node{map[string]bool{key: true}, cn.val + 1}, cur)
	} else {
		nxt.Value.(node).keys[key] = true
		t.nodes[key] = nxt
	}
	delete(cn.keys, key)
	if len(cn.keys) == 0 {
		t.Remove(cur)
	}
}

func (t *AllOne) Dec(key string) {
	cur := t.nodes[key]

	cn := cur.Value.(node)
	pre := cur.Prev()
	if cn.val == 1 {
		delete(t.nodes, key)
	} else if pre == nil || pre.Value.(node).val < cn.val-1 {
		t.nodes[key] = t.InsertBefore(node{map[string]bool{key: true}, cn.val - 1}, cur)
	} else {
		pre.Value.(node).keys[key] = true
		t.nodes[key] = pre
	}

	delete(cn.keys, key)
	if len(cn.keys) == 0 {
		t.Remove(cur)
	}
}

func (t *AllOne) GetMaxKey() string {
	if b := t.Back(); b != nil {
		for x, _ := range b.Value.(node).keys {
			return x
		}
	}
	return ""
}

func (t *AllOne) GetMinKey() string {
	if f := t.Front(); f != nil {
		for x, _ := range f.Value.(node).keys {
			return x
		}
	}
	return ""
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
