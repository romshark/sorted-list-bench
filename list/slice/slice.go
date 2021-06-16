package slice

type any = interface{}

// List manages a sorted slice
type List struct {
	sortFn  func(i, j any) bool
	s       []interface{}
	version uint64
}

func New(sortFn func(i, j any) bool) *List {
	return &List{
		sortFn: sortFn,
	}
}

func (l *List) Len() int        { return len(l.s) }
func (l *List) Version() uint64 { return l.version }

func (l *List) Delete(o any) bool {
	i := l.IndexOf(o)
	if i < 0 {
		return false
	}

	l.s = append(l.s[:i], l.s[i+1:]...)
	l.version++
	return true
}

func (l *List) Push(o any) {
	if len(l.s) < 1 {
		l.s = append(l.s, o)
		l.version++
		return
	}
	var i int
	for ix, v := range l.s {
		i = ix
		if !l.sortFn(v, o) {
			i--
			break
		}
	}
	i++
	if i >= len(l.s) {
		l.s = append(l.s, o)
	} else {
		l.s = append(l.s[:i+1], l.s[i:]...)
		l.s[i] = o
	}
	l.version++
}

func (l *List) IndexOf(o any) int {
	for i, v := range l.s {
		if v == o {
			return i
		}
	}
	return -1
}

func (l *List) Contains(o any) bool {
	return l.IndexOf(o) > -1
}

func (l *List) Scan(
	after any,
	fn func(any) bool,
) {
	ls := l.s
	if after != nil {
		for i, v := range ls {
			if v == after {
				ls = l.s[i+1:]
				goto SCAN
			}
		}
		// If search was unsuccessful then just return
		return
	}
SCAN:
	for _, v := range ls {
		if !fn(v) {
			return
		}
	}
}
