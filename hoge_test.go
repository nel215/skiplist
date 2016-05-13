package skiplist

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

func TestHoge(t *testing.T) {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	list := New(logger)
	list.Insert(6, struct{}{})
	list.Insert(10, struct{}{})
	list.Insert(8, struct{}{})
	list.Insert(12, struct{}{})
	iter := list.Iterator()
	for iter.Next() {
		logger.Println(iter.Key())
	}
}
