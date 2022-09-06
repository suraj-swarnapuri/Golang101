package main

import (
	"fmt"
)


func main(){
	/*
	Basic Types:
		- int, int8, int16, int32, int64
		- uint, uint8, uint16, uint32, uint64
		- float32, float64
		- complex64, complex128
		- byte -> uint8, rune -> int32
		- bool
		- string
	*/
	
	var a int = 10;
	println(a)

	var b bool = true;
	println(b)
	println("==========Type Inference==============")

	var c = 10;
	println(c)

	d := 10; // implicit type inference
	println(d)

	println("=========Constants================")
	const e = 10;
	// e = 20;
	println(e)

	println("=========Functions================")

	println(add(10,20))
	whole, remainder := splitDecimal(10.5)
	println(whole, remainder)
	
	// functions are values too!
	fAdd := add
	fMultiply := multiply
	println(foo(10,20,fAdd))
	println(foo(10,20,fMultiply))

	// functions as closures. Closures are functions that reference variables from outside their scope.
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}


	// defer: execute at the end of the function
	print("hello")
	defer print("world")
	println("!")

}

// Function Definition
func add(a , b int)int {
	return a+b
}
func multiply(a,b int)int{
	return a*b
}

func splitDecimal(a float64) (whole , remainder int) {
	// naked return
	whole = int(a)
	remainder = int(a*10)%10
	return
}
func foo(a,b int , f func(int,int)int)(int) {
	return f(a,b)
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first,second := 0,1
	return func()int {
	  c := first
	  first,second = second, first+second
	  return c
	  }
	
}