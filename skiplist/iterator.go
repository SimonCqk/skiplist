package skiplist

type skipListIterator struct {
	list *SkipList
	node *skipListNode
}

func (i *skipListIterator) Valid() bool {
	return i.node != nil
}

func (i *skipListIterator) Key() interface{} {
	return i.node.Key
}

func (i *skipListIterator) Next() {
	checkValid(i)
	i.node = i.node.Next()
}

func (i *skipListIterator) Prev() {
	checkValid(i)
	i.node = i.list.findLessThan(i.node.Key)
	if i.node == i.list.head {
		i.node = nil
	}
}

func (i *skipListIterator) Seek(target interface{}) {
	i.node = i.list.findGreaterOrEqual(target, nil)
}

func (i *skipListIterator) SeekToFirst() {
	i.node = i.list.head.Next()
}

func (i *skipListIterator) SeekToLast() {
	i.node = i.list.findLast()
	if i.node == i.list.head {
		i.node = nil
	}
}

func checkValid(it *skipListIterator) {
	if !it.Valid() {
		panic("SkipList iterator not valid!")
	}
}
