package slice_test

import (
	"testing"

	"github.com/romshark/sorted-list-bench/list/slice"

	"github.com/stretchr/testify/require"
)

func sortFunc(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func TestPushDelete(t *testing.T) {
	l := slice.New(sortFunc)
	Expect(t, l)
	require.Equal(t, uint64(0), l.Version())

	l.Push(1)
	Expect(t, l, 1)
	require.Equal(t, uint64(1), l.Version())

	l.Push(2)
	Expect(t, l, 1, 2)
	require.Equal(t, uint64(2), l.Version())

	l.Push(9)
	Expect(t, l, 1, 2, 9)
	require.Equal(t, uint64(3), l.Version())

	l.Push(6)
	l.Push(4)
	Expect(t, l, 1, 2, 4, 6, 9)
	require.Equal(t, uint64(5), l.Version())

	l.Push(0)
	Expect(t, l, 0, 1, 2, 4, 6, 9)
	require.Equal(t, uint64(6), l.Version())

	require.False(t, l.Delete(5))
	Expect(t, l, 0, 1, 2, 4, 6, 9)
	require.Equal(t, uint64(6), l.Version())

	require.True(t, l.Delete(6))
	Expect(t, l, 0, 1, 2, 4, 9)
	require.Equal(t, uint64(7), l.Version())

	require.True(t, l.Delete(0))
	Expect(t, l, 1, 2, 4, 9)
	require.Equal(t, uint64(8), l.Version())

	require.True(t, l.Delete(9))
	Expect(t, l, 1, 2, 4)
	require.Equal(t, uint64(9), l.Version())

	require.True(t, l.Delete(1))
	require.True(t, l.Delete(2))
	require.True(t, l.Delete(4))
	Expect(t, l)
	require.Equal(t, uint64(12), l.Version())
}

func TestScanInterrupt(t *testing.T) {
	l := slice.New(sortFunc)

	l.Push(1)
	l.Push(2)
	l.Push(9)
	l.Push(6)
	l.Push(4)
	l.Push(0)

	{
		called := 0
		l.Scan(nil, func(i interface{}) bool {
			called++
			return true
		})
		require.Equal(t, 6, called)
	}

	{
		called := 0
		l.Scan(nil, func(i interface{}) bool {
			called++
			return false
		})
		require.Equal(t, 1, called)
	}
}

func Expect(t *testing.T, l *slice.List, expected ...int) {
	require.Equal(t, len(expected), l.Len())
	actual := make([]int, 0, l.Len())
	l.Scan(nil, func(n interface{}) bool {
		actual = append(actual, n.(int))
		return true
	})
	if expected == nil {
		expected = []int{}
	}
	require.Equal(t, expected, actual)

	for _, e := range expected {
		require.True(t, l.Contains(e))
	}
}
