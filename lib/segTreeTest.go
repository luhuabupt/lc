package main

func main() {

}

func newSgTree(a []int) {
	t := make(sg, 4*len(a))
	t.build(a, 1, 1, len(a))
}

type sg []struct {
	l, r int
	val  int
}

// ((
func (t sg) op(x, y int) int {
	return x + y
}

func (t sg) set(o, v int) {
	t[o].val = v
}

// ))

func (t sg) maintain(o int) {
	ls, rs := t[o<<1], t[o<<1+1]
	t[o].val = t.op(ls.val, rs.val)
}

func (t sg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1+1, i, val)
	}
	t.maintain(o)
}

func (t sg) query(o, l, r int) int {
	if t[o].l >= l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1+1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1, l, r)
	return t.op(vl, vr)
}

func (t sg) queryAll() int {
	return t[1].val
}

func (t sg) build(a []int, o, l, r int) {
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1+1, m+1, r)
	t.maintain(o)
}
