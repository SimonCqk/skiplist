package skiplist

import "sync"

type skipListNode struct {
	Key  interface{}
	next *skipListNode
	sync.RWMutex
}

func (n *skipListNode) Next() *skipListNode {
	n.RLock()
	defer n.RUnlock()
	return n.next
}

func (n *skipListNode) SetNext(next *skipListNode) {
	n.Lock()
	n.next = next
	n.Unlock()
}
