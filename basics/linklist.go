// output
// 2--->3--->4--->5--->6--->
// 3--->4--->5--->6--->
// Empty Linked-list (underFlow)
// Empty Linked-list (no items)

// TODO: remove from end &  and remove at index ,
// sorting,searching
package main

import "fmt"

type node struct {
	value int
	next  *node
}

type List struct {
	first *node
	last  *node
	size  int
}

var list = &List{size: 0}

func append(value int) {
	newNode := node{value: value}
	if list.first == nil {
		list.first = &newNode
		list.last = &newNode
	} else {
		list.last.next = &newNode
		list.last = &newNode
	}
	list.size++
}

func prepend(value int) {
	newNode := node{value: value}
	if list.first == nil {
		list.first = &newNode
		list.last = &newNode

	} else {
		newNode.next = list.first
		list.first = &newNode
	}
	list.size++
}
func removeFromBegning() {
	if list.first == nil {
		fmt.Println("Empty Linked-list (underFlow)")
	} else {
		list.first = list.first.next
		list.size--
	}
}

func removeAtIndex(index int) {
	if list.first == nil {
		fmt.Println("Empty Linked-list (no items)")
		return
	}
	if index < 1 {
		fmt.Println("Invalid Entry")
		return
	}
	searchNode := list.first
	counter := 1
	for searchNode != nil {
		if counter == index-1 {
			fmt.Println("removed ", searchNode.next.value)
			searchNode.next = searchNode.next.next
			list.size--
			return
		}
		counter++
		searchNode = searchNode.next
	}
	fmt.Println("index out of bound ")

}

func removeFromLast() {
	if list.first == nil {
		fmt.Println("Empty Linked-list (underFlow)")
		return
	} else {
		removeAtIndex(list.size)
	}
}

func search(value int) {
	searchNode := list.first
	if list.first == nil {
		fmt.Println("Empty Linked-list (no items)")
		return
	}
	counter := 1
	for searchNode != nil {
		if searchNode.value == value {
			fmt.Println(value, " found at ", counter)
			return
		}
		counter++
		searchNode = searchNode.next
	}
	fmt.Println(value, "not found")
}

func swapNodes(x, y *node) {
	if x == y {
		return
	}

	var px, cx, py, cy *node

	cx = list.first
	for cx != nil && cx != x {
		px = cx
		cx = cx.next
	}

	cy = list.first
	for cy != nil && cy != y {
		py = cy
		cy = cy.next
	}

	// If either x or y is not present, nothing to do
	if cx == nil || cy == nil {
		return
	}
	// If x is not head of linked list
	if px != nil {
		px.next = cy
	} else {
		list.first = cy
	} // Else make y as new head

	// If y is not head of linked list
	if py != nil {
		py.next = cx
	} else // Else make x as new head
	{
		list.first = cx
	}

	// Swap next pointers
	temp := cy.next
	cy.next = cx.next
	cx.next = temp
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

func bubbleSort() {
	if list.first == nil {
		fmt.Println("Empty Linked-list (no items)")
		return
	}
	cy := list.first
	for i := 0; i < list.size; i++ {
		cy = list.first
		for cy != nil {
			if cy.next != nil && cy.value > cy.next.value {
				swapNodes(cy, cy.next)
			}
			cy = cy.next
		}
	}
}
func main() {

	append(4)
	append(5)
	removeAtIndex(5)
	fmt.Println(list.size)
	append(6)
	prepend(3)
	prepend(2)
	display()
	search(8)
	fmt.Println(list.size)
	removeAtIndex(5)
	removeFromLast()
	removeFromBegning()
	display()
	removeFromBegning()
	display()

	prepend(3)
	prepend(4)
	prepend(1)
	prepend(5)
	prepend(6)
	prepend(2)

	display()
	bubbleSort()

	display()

}
