package skiplist

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxHeight     = 16
	listBranching = 8
)

// Comparator compares values of two keys and return result by int
// -1: lhs < rhs / 0: lhs == rhs / 1: lhs > rhs
type Comparator func(lhs interface{}, rhs interface{}) int

type SkipList struct {
	maxHeight  int
	head       *skipListNode
	comparator Comparator
	sync.RWMutex
}

func newNode(key interface{}, height int) *skipListNode {
	return &skipListNode{
		Key:  key,
		next: make([]*skipListNode, height),
	}
}

func (l *SkipList) Init(comparator Comparator) {
	l.comparator = comparator
	l.head = newNode(0, 1)
}

func (l *SkipList) Contains(key interface{}) bool {
	n := l.findGreaterOrEqual(key, nil)
	return n != nil && l.comparator(key, n.Key) == 0
}

func (l *SkipList) Insert(key interface{}) {
	prev := make([]*skipListNode, maxHeight)
	n := l.findGreaterOrEqual(key, prev)
	if n != nil && l.comparator(key, n.Key) == 0 {
		panic("Insert duplicate key into SkipList")
	}
	height := l.randomHeight()
	if height > l.getMaxHeight() {
		for i := l.getMaxHeight(); i < height; i++ {
			prev[i] = l.head
		}
		l.maxHeight = height
	}
	n = newNode(key, height)
	for i := 0; i < height; i++ {
		n.SetNext(i, prev[i].Next(i))
		prev[i].SetNext(i, n)
	}
}

func (l *SkipList) keyIsAfterNode(key interface{}, node *skipListNode) bool {
	return node != nil && l.comparator(node.Key, key) < 0
}

func (l *SkipList) findGreaterOrEqual(key interface{}, prev []*skipListNode) *skipListNode {
	n := l.head
	level := l.getMaxHeight() - 1
	var next *skipListNode
	for {
		next = n.Next(level)
		if l.keyIsAfterNode(key, next) {
			n = next
		} else {
			if prev != nil {
				prev[level] = n
			}
			if level == 0 {
				return n
			} else {
				level--
			}
		}
	}
}

func (l *SkipList) findLessThan(key interface{}) *skipListNode {
	n := l.head
	level := l.getMaxHeight() - 1
	var next *skipListNode
	for {
		// assertion
		if !(n == l.head || l.comparator(n.Key, key) < 0) {
			panic(fmt.Sprintf("Invalid searching in findLessThan(%v)", key))
		}
		next = n.Next(level)
		if next == nil || l.comparator(next.Key, key) >= 0 {
			if level == 0 {
				return n
			} else {
				// switch to next list
				level--
			}
		} else {
			n = next
		}
	}
}

func (l *SkipList) findLast() *skipListNode {
	n := l.head
	level := l.getMaxHeight() - 1
	var next *skipListNode
	for {
		next = n.Next(level)
		if next == nil {
			if level == 0 {
				return n
			} else {
				// switch to next list
				level--
			}
		} else {
			n = next
		}
	}
}

func (l *SkipList) randomHeight() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	height := 1
	for height < maxHeight && r.Int()%listBranching == 0 {
		height++
	}
	return height
}

func (l *SkipList) getMaxHeight() int {
	l.RLock()
	defer l.RUnlock()
	return l.maxHeight
}
