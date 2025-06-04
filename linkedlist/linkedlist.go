package linkedlist

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList) Insert(value interface{}) {
	newNode := &Node{
		value: value,
		next:  l.head,
	}
	l.head = newNode
	l.size++
}

func (l *LinkedList) Delete(value interface{}) bool {
	if l.head == nil {
		return false
	}

	if l.head.value == value {
		l.head = l.head.next
		l.size--
		return true
	}

	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.size--
			return true
		}
		current = current.next
	}

	return false
}

func (l *LinkedList) Search(value interface{}) bool {
	current := l.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}
