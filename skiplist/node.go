package skiplist

import "sync"

type skipListNode struct {
	Key  interface{}
	next []*skipListNode
	sync.RWMutex
}

func (n *skipListNode) Next(level int) *skipListNode {
	n.RLock()
	defer n.RUnlock()
	return n.next[level]
}

func (n *skipListNode) SetNext(level int, next *skipListNode) {
	n.Lock()
	n.next[level] = next
	n.Unlock()
}
