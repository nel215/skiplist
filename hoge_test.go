package skiplist

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

func newList() *SkipList {
	logger := logrus.New()
	//logger.Level = logrus.DebugLevel
	logger.Level = logrus.WarnLevel
	return New(logger)
}

func TestHoge(t *testing.T) {
	list := newList()
	list.Insert(6, struct{}{})
	list.Insert(10, struct{}{})
	list.Insert(8, struct{}{})
	list.Insert(12, struct{}{})
	iter := list.Iterator()
	for iter.Next() {
		t.Log(iter.Key())
	}
}

func TestFind(t *testing.T) {
	list := newList()
	list.Insert(6, struct{}{})
	list.Insert(10, struct{}{})
	if _, ok := list.Find(6); !ok {
		t.Error("6 must be found")
	}
	if _, ok := list.Find(8); ok {
		t.Error("8 must not be found")
	}
	iter, _ := list.Find(6)
	iter.Next()
	if iter.Key() != 10 {
		t.Error("next key must be 10")
	}
}
