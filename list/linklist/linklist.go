package linklist

type any = interface{}

// List manages a sorted doubly-linked list
type List struct {
	sortFn func(i, j any) bool
	// prev maps values to predecessor nodes
	// or to the begining of the list defined by the actual *List
	prev      map[any]any
	firstNode *node
	len       int
	version   uint64
}

func New(sortFn func(i, j any) bool) *List {
	return &List{
		sortFn: sortFn,
		prev:   map[any]any{},
	}
}

func (l *List) Len() int        { return l.len }
func (l *List) Version() uint64 { return l.version }

func (l *List) Delete(i any) bool {
	p, ok := l.prev[i]
	if !ok {
		return false
	}

	switch p := p.(type) {
	case *List:
		// At head
		l.firstNode = l.firstNode.Next
		if l.firstNode != nil {
			l.prev[l.firstNode.Data] = l
		}
	case *node:
		// At body
		if p.Next.Next != nil {
			// Next has a follower
			l.prev[p.Next.Next.Data] = p
		}
		p.Next = p.Next.Next

	}
	delete(l.prev, i)
	l.len--
	l.version++
	return true
}

func (l *List) Push(o any) {
	onNew := func(newNode *node, previous any) {
		l.prev[o] = previous
		l.len++
		l.version++
	}
	if l.firstNode == nil {
		// At head of empty list
		l.firstNode = &node{Data: o}
		onNew(l.firstNode, l /* Point to the list itself */)
		return

	}
	var previous *node
	current := l.firstNode
	for {
		if l.sortFn(current.Data, o) {
			if current.Next == nil {
				// At tail
				current.Next = &node{Data: o}
				onNew(current.Next, current)
				return
			}
			// Move to next node
			previous = current
			current = current.Next
			continue
		} else {
			if previous == nil {
				// At head
				prevData := l.firstNode.Data
				l.firstNode = &node{Data: o, Next: l.firstNode}
				l.prev[prevData] = l.firstNode // Update previous
				onNew(l.firstNode, l /* Point to the list itself */)
				return
			}
			// At body
			previous.Next = &node{Data: o, Next: current}
			l.prev[current.Data] = previous.Next
			onNew(previous.Next, previous)
			return
		}
	}
}

func (l *List) Contains(o any) bool {
	_, ok := l.prev[o]
	return ok
}

func (l *List) getNodeByData(data any) *node {
	if x, ok := l.prev[data]; ok {
		switch x := x.(type) {
		case *List:
			return x.firstNode
		case *node:
			return x
		}
	}
	return nil
}

func (l *List) Scan(
	after any,
	fn func(any) bool,
) {
	current := l.firstNode
	if after != nil {
		if n := l.getNodeByData(after); n == nil {
			return
		} else {
			current = n
		}
	}
	for current != nil {
		if !fn(current.Data) {
			return
		}
		current = current.Next
	}
}

type node struct {
	Data any
	Next *node
}
