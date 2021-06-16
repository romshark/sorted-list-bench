package list

import (
	"github.com/romshark/sorted-list-bench/list/linklist"
	"github.com/romshark/sorted-list-bench/list/skiplist"
	"github.com/romshark/sorted-list-bench/list/slice"
)

type List interface {
	Len() int
	Version() uint64
	Delete(o interface{}) bool
	Push(o interface{})
	Contains(o interface{}) bool
	Scan(after interface{}, fn func(interface{}) bool)
}

type Implementation struct {
	Name string
	Make func() List
}

func Implementations(sortFn func(i, j interface{}) bool) []Implementation {
	return []Implementation{
		{"slice", func() List { return slice.New(sortFn) }},
		{"linklist", func() List { return linklist.New(sortFn) }},
		{"skiplist", func() List { return skiplist.New() }},
	}
}
