package skiplist

type SkipList struct {
	maxHeight int
	head      *skipListNode
}

func (l *SkipList) Contains(key interface{}) bool {
	return false
}
