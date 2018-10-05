package skiplist

type SkipList struct {
	maxHeight int
	head      *skipListNode
}

func newNode(key interface{}, height int) {

}

func (l *SkipList) Contains(key interface{}) bool {
	return false
}

func (l *SkipList) Insert(key interface{}) {

}

func (l *SkipList) keyIsAfterNode(key interface{}, node *skipListNode) bool {
	return false
}

func (l *SkipList) findGreaterOrEqual(key interface{}, node **skipListNode) *skipListNode {
	return nil
}

func (l *SkipList) findLessThan(key interface{}) *skipListNode {
	return nil
}

func (l *SkipList) findLast() *skipListNode {
	return nil
}

func (l *SkipList) randomHeight() int {
	return -1
}
