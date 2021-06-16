package list

import (
	"github.com/romshark/sorted-list-bench/list/linklist"
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
		{"linklist", func() List { return linklist.New(sortFn) }},
		{"slice", func() List { return slice.New(sortFn) }},
	}
}
