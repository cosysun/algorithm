package linklist

import (
	"fmt"
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

func TestSingleLinkList_HasLoop(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	//list.GenerateLoopList(10)
	//	list.Print()
	isLoop := list.HasLoop()
	if isLoop {
		fmt.Println("has loop in link list")
	} else {
		fmt.Println("no loop in link list")
	}
}

func TestMergeTwoLinkList(t *testing.T) {
	list1 := NewSingleLink()
	list1.InsertToTail(1)
	list1.InsertToTail(4)
	list1.InsertToTail(9)
	list1.Print()

	list2 := NewSingleLink()
	for i := 0; i < 10; i++ {
		list2.InsertToTail(i)
	}
	list2.Print()

	list := MergeTwoLinkList(list1, list2)
	list.Print()

}

func TestMergeTwoLinkList2(t *testing.T) {
	list1 := NewSingleLink()
	list1.InsertToTail(1)
	list1.InsertToTail(4)
	list1.InsertToTail(9)
	list1.Print()

	list2 := NewSingleLink()
	for i := 0; i < 10; i++ {
		list2.InsertToTail(i)
	}
	list2.Print()

	list := NewSingleLink()
	list.head.next = MergeTwoLinkList2(list1.head.next, list2.head.next)
	list.Print()
}

func TestSingleLinkList_DeleteReverseNode(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	list.DeleteReverseNode(7)
	list.Print()
}

func TestSingleLinkList_DeleteReverseNode1(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	list.DeleteReverseNode1(7)
	list.Print()
}

func TestSingleLinkList_FindMiddleNode(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	fmt.Println(list.FindMiddleNode().GetValue().(int))
}

func TestSingleLinkList_FindMiddleNode1(t *testing.T) {
	list := NewSingleLink()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	fmt.Println(list.FindMiddleNode1().GetValue().(int))
}
