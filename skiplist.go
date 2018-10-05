package skiplist

import "sync"

type skipListNode struct {
	Key  interface{}
	Next *skipListNode
	sync.RWMutex
}

type skipListIterator struct {
	list *SkipList
	node *skipListNode
}

type SkipList struct {
	maxHeight int
	head      *skipListNode

	node struct {
		key interface{}
	}
}

func (i *skipListIterator) Valid() bool {
	return i.node != nil
}

func (i *skipListIterator) Key() interface{} {
	return nil
}

func (i *skipListIterator) Next() {

}

func (i *skipListIterator) Prev() {

}

func (i *skipListIterator) Seek() {

}

func (i *skipListIterator) SeekToFirst() {

}

func (i *skipListIterator) SeekToLast() {

}

func (l *SkipList) Contains(key interface{}) bool {

}
