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

}

func (i *skipListIterator) Seek() {

}

func (i *skipListIterator) SeekToFirst() {

}

func (i *skipListIterator) SeekToLast() {

}

func checkValid(iter *skipListIterator) {
	if !iter.Valid() {
		panic("SkipList iterator not valid!")
	}
}
