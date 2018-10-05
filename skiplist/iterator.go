package skiplist

type skipListIterator struct {
	list *SkipList
	node *skipListNode
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
