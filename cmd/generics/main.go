package main

import (
	"fmt"
)

func Index[T comparable](s []T, x T)int{
	for i, v := range s{
		if v == x{
			return i
		}
	}
	return -1
}

type Node[T comparable] struct {
	next *Node[T]
	val  T
}
func (n *Node[T]) PrintList(){
	for{
		if n != nil{
			println(n.val)
		}else{
			return
		}
		n = n.next
	}
}

func main(){
	fmt.Println("=========Generic Functions================")

	fmt.Println(Index([]int{1,2,3,4,5}, 3))
	fmt.Println(Index([]string{"a","b","c","d","e"}, "b"))
	fmt.Println("=========Generic Types================")
	
	in := &Node[int]{next: nil, val: 1}
	in.next = &Node[int]{next: nil, val: 2}
	in.next.next = &Node[int]{next: nil, val: 3}
	
	in.PrintList()

	sn := &Node[string]{next: nil, val: "a"}
	sn.next = &Node[string]{next: nil, val: "b"}
	sn.next.next = &Node[string]{next: nil, val: "c"}
	sn.PrintList()

}