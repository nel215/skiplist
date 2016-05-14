package skiplist

import (
	"github.com/Sirupsen/logrus"
)

type Iterator struct {
	node   *node
	logger logrus.FieldLogger
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
