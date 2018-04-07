package datastructure

type Node struct {
	Value string
	Next  *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) isEmpty() bool {
	if q.head == nil && q.tail == nil {
		return true
	}
	return false
}

func (q *Queue) Add(str string) {
	n := &Node{str, nil}
	if q.tail == nil {
		q.tail = n
	} else {
		q.tail.Next = n
		if q.head == nil {
			q.head = q.tail
		}
		q.tail = n
	}
}

func (q *Queue) Remove() (string, bool) {
	if q.isEmpty() {
		return "", false
	}
	if q.head == nil && q.tail != nil {
		str := q.tail.Value
		q.tail = nil
		return str, true
	}
	if q.head.Next == q.tail {
		str := q.head.Value
		q.head = nil
		return str, true
	}
	n := q.head.Next
	str := q.head.Value
	q.head = n
	return str, true
}

func (q Queue) String() string {
	if q.isEmpty() {
		return ""
	}
	if q.head == nil {
		return q.tail.Value
	}
	n := q.head
	str := n.Value
	for n.Next != nil {
		n = n.Next
		str += (" - " + n.Value)
	}
	return str
}
