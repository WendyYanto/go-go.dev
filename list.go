package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (node *List[int]) GetCount() int32 {
	result := 0

	curr := node

	for curr != nil {
		result++
		curr = curr.next
	}

	return int32(result)
}

func main() {
	head := List[int] {
		next: &List[int] {
			next: &List[int] {
				next: nil,
				val: 3,
			},
			val: 2,
		},
		val: 1,
	}

	fmt.Println(head.GetCount())
}
