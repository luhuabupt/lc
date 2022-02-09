package classic

type node struct {
	next *node
	pre  *node
	k    int
}

type LFUCache struct {
	cap     int
	kv      map[int]int
	cnt     map[int]*node
	list    map[int]*node
	minMap  map[int]int
	minHead *node
	minTail *node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		kv:      nil,
		cnt:     nil,
		list:    nil,
		minMap:  nil,
		minHead: nil,
		minTail: nil,
	}
}

func (this *LFUCache) Get(key int) int {
	if val, ok := this.kv[key]; ok {
		this.counter(key)
		return val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if _, ok := this.kv[key]; !ok && len(this.cnt) == this.cap {
		this.del()
	}

	this.counter(key)
	this.kv[key] = value
}

func (this *LFUCache) counter(key int) {

}

func (this *LFUCache) del() {
	tail := this.minTail
	if this.minMap[tail.k] == 1 {
		this.minTail = this.minTail.pre
		delete(this.minMap, tail.k)

	} else {

	}

}
