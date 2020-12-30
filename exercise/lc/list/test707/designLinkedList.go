package test707

type MyLinkedList struct {
	dummy *Node
	size  int
}
type Node struct {
	val  int
	next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &Node{},
		size:  0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (list *MyLinkedList) Get(index int) int {
	if index >= list.size || index < 0 {
		return -1
	}
	cur := list.dummy
	for i := 0; i < index+1; i++ {
		cur = cur.next
	}
	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (list *MyLinkedList) AddAtHead(val int) {
	list.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (list *MyLinkedList) AddAtTail(val int) {
	list.AddAtIndex(list.size, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index > list.size {
		return
	}
	if index < 0 {
		index = 0
	}
	list.size++
	prev := list.dummy
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &Node{val, prev.next}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (list *MyLinkedList) DeleteAtIndex(index int) {

	if index < 0 || list.size <= index {
		return
	}
	prev := list.dummy
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	ret := prev.next
	prev.next = ret.next
	list.size--
}
