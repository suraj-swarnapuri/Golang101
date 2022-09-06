package main

import (
	"fmt"
)

type calc interface{
	add(int, int) int
	sub(int, int) int
	mul(int, int) int
	div(int, int) int
}

type Calc struct{}

func (c Calc) add(a,b int)int{
	return a+b
}

func (c Calc) sub(a,b int)int{
	return a-b
}

func (c Calc) mul(a,b int)int{
	return a*b
}

func (c Calc) div(a,b int)(int,error){
	if b == 0{
		return 0, fmt.Errorf("Division by zero")
	}
	return a/b, nil
}