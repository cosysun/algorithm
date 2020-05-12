package linklist

import (
	"testing"
)

func TestSingleLinkList_InsertToHead(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 10; i++ {
		list.InsertToHead(i)
	}
	list.Print()
}

func TestSingleLinkList_InsertToTail(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
}

func TestSingleLinkList_Reverse(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	list.Reverse()
	list.Print()
}

func TestReverse2(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.head.next = Reverse2(list.head.next)
	list.Print()
}
