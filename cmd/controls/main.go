package main

func main(){
	println("=========For Loop================")

	for i := 0; i < 10; i++ {
		println(i)
	}
	
	println("=========While Loop================")
	// For is Gos's while loop
	sum := 0 
	for sum < 10 {
		sum += 1
	}
	println(sum)
	println("==========If Statements==============")

	a := 10
	if a < 10 {
		println("a is less than 10")
	} else if a == 10 {
		println("a is equal to 10")
	} else {
		println("a is greater than 10")
	}

	if b:=4; b < 10 {
		println("b is less than 10")
	}
	println("==========Switch Statements==============")
	str := "Hello"

	switch str {
	case "Hello":
		println("Hello")
	case "Hi":
		println("Hi")
	default:
		println("Default")
	}
	// alternative to if-elseif-else chains
	a = 10
	switch {
	case a < 10:
		println("a is less than 10")
	case a == 10:
		println("a is equal to 10")
	case a > 10:
		println("a is greater than 10")
	}
}