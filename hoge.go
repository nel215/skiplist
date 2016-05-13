package skiplist

import (
	_ "log"
)

type min struct{}
type max struct{}

type node struct {
	key  interface{}
	next *node
}

type SkipList struct {
	head *node
}

func New() *SkipList {
	tail := &node{max{}, nil}
	head := &node{min{}, tail}
	return &SkipList{head}
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

func (s *SkipList) Insert(key interface{}, value interface{}) {
	e := &node{key, nil}
	now := s.head
	for compare(now.next.key, e.key) {
		now = now.next
	}
	e.next = now.next
	now.next = e
}

func (s *SkipList) Iterator() *Iterator {
	return &Iterator{s.head}
}

type Iterator struct {
	node *node
}

func (i *Iterator) Next() bool {
	if _, ok := i.node.next.key.(max); ok {
		return false
	}
	i.node = i.node.next
	return true
}

func (i *Iterator) Key() interface{} {
	return i.node.key
}
