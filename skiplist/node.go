package skiplist

import "sync"

type skipListNode struct {
	Key  interface{}
	Next *skipListNode
	sync.RWMutex
}
