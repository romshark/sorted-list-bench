package llrb

import (
	"fmt"

	"github.com/petar/GoLLRB/llrb"
)

type Item llrb.Item

type Valuer interface {
	Value() interface{}
}

type ListItem interface {
	Item
	Valuer
}

func New(itemFactory func(interface{}) ListItem) *List {
	return &List{
		t:    llrb.New(),
		fact: itemFactory,
	}
}

type List struct {
	t       *llrb.LLRB
	fact    func(interface{}) ListItem
	version uint64
}

func (l *List) Len() int {
	return l.t.Len()
}

func (l *List) Version() uint64 {
	return l.version
}

func (l *List) Delete(o interface{}) bool {
	fmt.Println(o)
	if e := l.t.Delete(l.fact(o)); e == nil {
		return false
	}
	l.version++
	return true
}

func (l *List) Push(o interface{}) {
	l.t.InsertNoReplace(l.fact(o))
	l.version++
}

func (l *List) Contains(o interface{}) bool {
	return l.t.Get(l.fact(o)) != nil
}

func (l *List) Scan(after interface{}, fn func(interface{}) bool) {
	var pivot llrb.Item
	if after == nil {
		pivot = l.t.Min()
	} else {
		pivot = l.fact(after)
	}
	l.t.AscendGreaterOrEqual(pivot, func(i llrb.Item) bool {
		ai := i.(ListItem)
		return fn(ai.Value())
	})
}
