package skiplist

import (
	"github.com/Sirupsen/logrus"
	"math/rand"
)

const (
	MAX_DEPTH = 20
)

type min struct{}
type max struct{}

type node struct {
	key  interface{}
	next *node
	down *node
}

type SkipList struct {
	heads  []*node
	logger logrus.FieldLogger
}

func New(logger logrus.FieldLogger) *SkipList {
	s := &SkipList{
		heads:  make([]*node, MAX_DEPTH),
		logger: logger,
	}
	for i := 0; i < MAX_DEPTH; i++ {
		tail := &node{max{}, nil, nil}
		head := &node{min{}, tail, nil}
		s.heads[i] = head
	}
	for i := 0; i < MAX_DEPTH-1; i++ {
		s.heads[i].down = s.heads[i+1]
	}
	return s
}

func (s *SkipList) bottom() *node {
	return s.heads[MAX_DEPTH-1]
}

func compare(a, b interface{}) bool {
	if _, ok := a.(min); ok {
		return true
	}
	if _, ok := a.(max); ok {
		return false
	}
	switch v := a.(type) {
	case int:
		return v < b.(int)
	default:
		return false
	}
}

func (s *SkipList) insert(now *node, depth int, key interface{}) *node {
	if depth == MAX_DEPTH {
		return nil
	}
	s.logger.Debug(depth, now.key)
	for compare(now.next.key, key) {
		now = now.next
	}

	s.logger.Debug(now.down, now.key)
	down := s.insert(now.down, depth+1, key)

	if depth == MAX_DEPTH-1 || (down != nil && rand.Float64() < 0.5) {
		e := &node{
			key:  key,
			next: now.next,
			down: down,
		}
		now.next = e
		return e
	}
	return nil
}

func (s *SkipList) Insert(key interface{}, value interface{}) {
	s.insert(s.heads[0], 0, key)
}

func (s *SkipList) remove(now *node, depth int, key interface{}) {
	if depth == MAX_DEPTH {
		return
	}

	for compare(now.next.key, key) {
		now = now.next
	}
	s.remove(now.down, depth+1, key)
	for now.next.key == key {
		tmp := now.next.next
		now.next.down = nil
		now.next.next = nil
		now.next = tmp
	}
}

func (s *SkipList) Remove(key interface{}) {
	s.remove(s.heads[0], 0, key)
}

func (s *SkipList) find(now *node, depth int, key interface{}) *node {
	for compare(now.next.key, key) {
		now = now.next
	}
	if depth == MAX_DEPTH-1 {
		return now
	}
	return s.find(now.down, depth+1, key)
}
func (s *SkipList) Find(key interface{}) (*Iterator, bool) {
	n := s.find(s.heads[0], 0, key)
	if n.next.key == key {
		return &Iterator{n.next, s.logger}, true
	}
	return nil, false
}

func (s *SkipList) Iterator() *Iterator {
	return &Iterator{s.bottom(), s.logger}
}

type Iterator struct {
	node   *node
	logger logrus.FieldLogger
}

func (i *Iterator) Next() bool {
	a, b := i.node.next.key.(max)
	i.logger.Debug(a, b)
	if _, ok := i.node.next.key.(max); ok {
		return false
	}
	i.node = i.node.next
	return true
}

func (i *Iterator) Key() interface{} {
	return i.node.key
}
