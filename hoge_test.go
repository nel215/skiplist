package skiplist

import (
	"log"
	"testing"
)

func TestHoge(t *testing.T) {
	list := New()
	list.Insert(10, struct{}{})
	list.Insert(8, struct{}{})
	list.Insert(12, struct{}{})
	iter := list.Iterator()
	for iter.Next() {
		log.Println(iter.Key())
	}
}
