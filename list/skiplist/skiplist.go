package skiplist

import (
	"github.com/huandu/skiplist"
)

func New() *List {
	return &List{
		l: skiplist.New(skiplist.IntAsc),
	}
}

type List struct {
	l       *skiplist.SkipList
	version uint64
}

func (l *List) Len() int {
	return l.l.Len()
}

func (l *List) Version() uint64 {
	return l.version
}

func (l *List) Delete(o interface{}) bool {
	if e := l.l.Remove(o); e == nil {
		return false
	}
	l.version++
	return true
}

func (l *List) Push(o interface{}) {
	l.l.Set(o, o)
	l.version++
}

func (l *List) Contains(o interface{}) bool {
	_, ok := l.l.GetValue(o)
	return ok
}

func (l *List) Scan(after interface{}, fn func(interface{}) bool) {
	var e *skiplist.Element
	if after != nil {
		if e = l.l.Get(after); e == nil {
			return
		}
	} else {
		e = l.l.Front()
	}
	for ; e != nil; e = e.Next() {
		if !fn(e.Value) {
			return
		}
	}
}
