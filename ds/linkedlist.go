package ds

// Node is a node in a singly linked list.
type Node struct {
	Val  int
	Next *Node
}

// LinkedList is a singly linked list.
type LinkedList struct {
	head *Node
	size int
}

// NewLinkedList creates a new empty LinkedList.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// InsertAtHead inserts a value at the head of the list.
func (ll *LinkedList) InsertAtHead(val int) {
	ll.head = &Node{Val: val, Next: ll.head}
	ll.size++
}

// InsertAtTail inserts a value at the tail of the list.
func (ll *LinkedList) InsertAtTail(val int) {
	ll.size++
	if ll.head == nil {
		ll.head = &Node{Val: val}
		return
	}
	cur := ll.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{Val: val}
}

// Delete removes the first occurrence of val from the list.
// Returns true if val was found and removed.
func (ll *LinkedList) Delete(val int) bool {
	if ll.head == nil {
		return false
	}
	if ll.head.Val == val {
		ll.head = ll.head.Next
		ll.size--
		return true
	}
	cur := ll.head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			ll.size--
			return true
		}
		cur = cur.Next
	}
	return false
}

// Search returns true if val exists in the list.
func (ll *LinkedList) Search(val int) bool {
	cur := ll.head
	for cur != nil {
		if cur.Val == val {
			return true
		}
		cur = cur.Next
	}
	return false
}

// ToSlice returns the list contents as a slice (for testing).
func (ll *LinkedList) ToSlice() []int {
	result := make([]int, 0, ll.size)
	cur := ll.head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

// IsEmpty returns true if the list contains no elements.
func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

// Size returns the number of elements in the list.
func (ll *LinkedList) Size() int {
	return ll.size
}
