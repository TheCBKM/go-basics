// output
// 2--->3--->4--->5--->6--->
// 3--->4--->5--->6--->
// Empty Linked-list (underFlow)
// Empty Linked-list (no items)
package main

import "fmt"

type node struct {
	value int
	next  *node
}

type List struct {
	first *node
	last  *node
}

var list = &List{}

func append(value int) {
	newNode := node{value: value}
	if list.first == nil {
		list.first = &newNode
		list.last = &newNode
		return
	} else {
		list.last.next = &newNode
		list.last = &newNode
	}
}

func prepend(value int) {
	newNode := node{value: value}
	if list.first == nil {
		list.first = &newNode
		list.last = &newNode
		return
	} else {
		newNode.next = list.first
		list.first = &newNode
	}
}
func removeFromBegning() {
	if list.first == nil {
		fmt.Println("Empty Linked-list (underFlow)")
		return
	} else {
		list.first = list.first.next
	}
}

func display() {
	if list.first == nil {
		fmt.Println("Empty Linked-list (no items)")
		return
	}
	displayNode := list.first
	for displayNode != nil {
		fmt.Print(displayNode.value, "--->")
		displayNode = displayNode.next
	}
	fmt.Println()

}

func main() {
	append(4)
	append(5)
	append(6)
	prepend(3)
	prepend(2)
	display()
	removeFromBegning()
	display()
	removeFromBegning()
	removeFromBegning()
	removeFromBegning()
	removeFromBegning()
	removeFromBegning()

	display()

}
