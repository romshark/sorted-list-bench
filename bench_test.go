package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/romshark/sorted-list-bench/list"
	"github.com/romshark/sorted-list-bench/rnd"
)

func init() {
	rand.Seed(6532)
}

type List interface {
	Len() int
	Version() uint64
	Delete(o interface{}) bool
	Push(o interface{})
	Contains(o interface{}) bool
	Scan(after interface{}, fn func(interface{}) bool)
}

func BenchmarkPush(b *testing.B) {
	for _, s := range []struct {
		minValue int
		maxValue int
		listSize int
	}{
		{minValue: 1, maxValue: 1_000_000_000, listSize: 100},
		{minValue: 1, maxValue: 1_000, listSize: 100},
	} {
		b.Run(fmt.Sprintf(
			"sz=%d_min=%d_max=%d", s.listSize, s.minValue, s.maxValue,
		), func(b *testing.B) {
			v := rnd.MakeInts(s.listSize, s.minValue, s.maxValue)
			b.Log("generated random data set")

			for _, i := range list.Implementations(sort) {
				b.Run(i.Name, func(b *testing.B) {
					l := i.Make()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						for _, v := range v {
							l.Push(v)
						}
					}
				})
			}
		})
	}
}

func sort(i, j interface{}) bool {
	return i.(int) > j.(int)
}
