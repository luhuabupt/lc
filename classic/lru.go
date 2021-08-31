package classic

type LRUCache struct {
	c, l       int
	m          map[int]*Dlink
	head, tail *Dlink
}

type Dlink struct {
	k, v      int
	pre, next *Dlink
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		c:    capacity,
		m:    map[int]*Dlink{},
		head: &Dlink{},
		tail: &Dlink{},
	}
	lru.head.next, lru.tail.pre = lru.tail, lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if this.m[key] == nil {
		return -1
	}
	this.Del(this.m[key])
	this.MvToHead(this.m[key])
	return this.m[key].v
}

func (this *LRUCache) Put(key int, value int) {
	if this.m[key] != nil {
		this.Del(this.m[key])
		this.MvToHead(this.m[key])
		this.m[key].v = value
	} else {
		n := &Dlink{
			k: key,
			v: value,
		}
		this.m[key] = n
		this.MvToHead(n)

		if this.l == this.c { // 超出容量
			delete(this.m, this.tail.pre.k)
			this.Del(this.tail.pre)
		} else {
			this.l++
		}
	}
}

func (this *LRUCache) Del(cur *Dlink) {
	cur.pre.next, cur.next.pre = cur.next, cur.pre
}

func (this *LRUCache) MvToHead(cur *Dlink) {
	cur.next, this.head.next.pre, this.head.next, cur.pre = this.head.next, cur, cur, this.head
}
