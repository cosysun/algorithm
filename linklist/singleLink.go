package linklist

import "fmt"

type ListNode struct {
	value interface{}
	next  *ListNode
}

type SingleLinkList struct {
	head *ListNode
	len  uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{next: nil, value: v}
}

func (n *ListNode) GetValue() interface{} {
	return n.value
}

func (n *ListNode) GetNext() interface{} {
	return n.next
}

func NewSingleLink() *SingleLinkList {
	return &SingleLinkList{head: NewListNode(0), len: 0}
}

func (s *SingleLinkList) InsertAfter(node *ListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	vNode := NewListNode(v)
	vNode.next = node.next
	node.next = vNode
	s.len++
	return true
}

func (s *SingleLinkList) InsertBefore(node *ListNode, v interface{}) bool {
	if node == nil || node == s.head {
		return false
	}
	cur := s.head
	for cur.next != node {
		if cur.next == nil {
			return false
		}
		cur = cur.next
	}
	vNode := NewListNode(v)
	cur.next = vNode
	vNode.next = node
	s.len++
	return true
}

func (s *SingleLinkList) InsertToTail(v interface{}) bool {
	cur := s.head
	for nil != cur.next {
		cur = cur.next
	}
	return s.InsertAfter(cur, v)
}

func (s *SingleLinkList) InsertToHead(v interface{}) bool {
	return s.InsertAfter(s.head, v)
}

func (s *SingleLinkList) FindByIndex(index uint) *ListNode {
	if index > s.len {
		return nil
	}
	cur := s.head
	var i uint = 0
	for nil != cur {
		if i == index {
			return cur
		}
		cur = cur.next
		i++
	}
	return cur
}

func (s *SingleLinkList) DeleteNode(node *ListNode) bool {
	if node == nil || s.head == node {
		return false
	}

	cur := s.head
	for cur.next != node {
		if cur.next == nil {
			return false
		}
		cur = cur.next
	}
	cur.next = node.next
	node = nil
	s.len--
	return true
}

func (s *SingleLinkList) Print() {
	cur := s.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (s *SingleLinkList) Reverse() bool {
	if s.head.next == nil {
		return true
	}

	cur := s.head.next
	p := cur.next
	cur.next = nil
	s.head.next = cur
	for nil != p {
		tmp := p.next
		tmp1 := s.head.next
		s.head.next = p
		p.next = tmp1
		p = tmp
	}
	return true
}

func Reverse2(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	newHead := Reverse2(head.next)
	head.next.next = head
	head.next = nil
	return newHead

}
