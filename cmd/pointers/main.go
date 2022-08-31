package main

import (
	"fmt"
)
func main(){
	fmt.Println("=========Pointers================")
	i := 10 
	p := &i
	fmt.Println(i)
	fmt.Println(p)

	*p = 20 // dereference operator
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println("=========Structs================")
	type Point struct{
		X int
		Y int
	}
	p1 := Point{1,2}
	fmt.Println(&p1)
	fmt.Println("=========Arrays================")

	a := [8]int{1,2,3,4,5,6,7,8}
	fmt.Println(a)
	fmt.Println("=========Slices================")

	s := a[1:4]
	fmt.Println(s)
	// slices are references to the underlying array
	s[0] = 10
	fmt.Println(a)
	fmt.Println(s)
	// create a slice with make
	s2 := make([]int, 5)
	fmt.Println(s2)
	// append 
	s2 = append(s2, 1,2,3,4,5)
	fmt.Println(s2)
	// range
	for i, v := range s2 {
		// where i is the index and v is the value
		fmt.Println(i,v)
	}
	fmt.Println("=========Maps================")
	m := make(map[string]int) // create a nil map
	m["a"] = 1 // insert
	fmt.Println(m["a"]) // lookup
	delete(m,"a") // delete
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("a not found")
	}
	
}
