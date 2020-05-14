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

func (s *SingleLinkList) GenerateLoopList(index uint) {
	node := s.FindByIndex(index)
	cur := s.head
	for nil != cur.next {
		cur = cur.next
	}
	cur.next = node
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

func (s *SingleLinkList) HasLoop() bool {
	if s.head == nil || s.head.next == nil {
		return false
	}

	pFast := s.head.next
	pSlow := s.head.next

	for pFast.next != nil {
		pFast = pFast.next.next
		if pFast == nil {
			break
		}
		pSlow = pSlow.next
		if pSlow == nil {
			break
		}
		if pSlow == pFast {
			return true
		}
	}
	return false
}

func (s *SingleLinkList) DeleteReverseNode(index uint) bool {
	if index > s.len {
		return false
	}

	tmpIndex := s.len - index - 1
	cur := s.head.next
	var i uint = 0
	for nil != cur {
		if i == tmpIndex {
			tmp := cur.next
			if tmp == nil {
				return false
			}
			cur.next = tmp.next
			tmp.next = nil
		}
		cur = cur.next
		i++
	}
	//node := s.FindByIndex(tmpIndex)
	//s.DeleteNode(node)
	return true
}

func (s *SingleLinkList) DeleteReverseNode1(n uint) bool {
	pFirst := s.head.next
	pSecond := s.head.next
	var index uint = 0
	for pFirst != nil {
		pFirst = pFirst.next
		index++
		if index > n+1 {
			pSecond = pSecond.next
		}
	}
	if index < n {
		return false
	}
	tmp := pSecond.next
	pSecond.next = tmp.next
	tmp.next = nil
	return true
}

func (s *SingleLinkList) FindMiddleNode() *ListNode {
	index := (s.len + 1) / 2
	return s.FindByIndex(index)
}

func (s *SingleLinkList) FindMiddleNode1() *ListNode {
	pFast := s.head.next
	pSlow := s.head.next
	for nil != pFast && pFast.next != nil {
		pFast = pFast.next.next
		pSlow = pSlow.next
	}
	return pSlow
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

func MergeTwoLinkList(l1 *SingleLinkList, l2 *SingleLinkList) *SingleLinkList {
	if l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2.head == nil || l2.head.next == nil {
		return l1
	}
	cur1 := l1.head.next
	cur2 := l2.head.next
	result := NewSingleLink()
	n := result.head
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) <= cur2.value.(int) {
			n.next = cur1
			cur1 = cur1.next
		} else {
			n.next = cur2
			cur2 = cur2.next
		}
		result.len++
		n = n.next
	}
	if cur1 == nil {
		n.next = cur2
	}

	if cur2 == nil {
		n.next = cur1
	}
	return result
}

func MergeTwoLinkList2(n1 *ListNode, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n1
	}

	if n1.value.(int) <= n2.value.(int) {
		n1.next = MergeTwoLinkList2(n1.next, n2)
		return n1
	} else {
		n2.next = MergeTwoLinkList2(n1, n2.next)
		return n2
	}
}
