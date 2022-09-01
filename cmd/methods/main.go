package main

import (
	"fmt"
	"time"
)

type Human interface{
	greet() string
}

type Person struct{
	name string
}


func (p Person) greet() string{
	return fmt.Sprintf("Hello my name is %s", p.name)
}
func doGreeting(h Human){
	fmt.Println(h.greet())
}


type Developer struct{
	name string
	favLanguage string
}

type DevError struct{
	When time.Time
	What string
}
func (e *DevError) Error() string{
	return fmt.Sprintf("%s: %s", e.What, e.When)
}

func (d Developer) greet()string{
	// Methods are functions with a receiver argument.
	d.name = "suraj"
	return fmt.Sprintf("fmt.Println(\"Hello my name is %s\")", d.name)
} 
func (d *Developer) setFavoritelanguage(language string)error{
	if language != "Go"{
		return &DevError{time.Now(), "Go is the only language I know"}
	}
	d.favLanguage = language
	return nil
}

func (d Developer) returnName() string{
	return d.name
}

func (d *Developer) changeName(newName string){
	// pointer receivers are used to modify the value of the receiver
	d.name = newName
}

func main(){
	fmt.Println("===Interfaces===")

	// Interfaces are implemented implicitly
	p := Person{name:"Chris"}
	doGreeting(p)
	// The Empty interface specifies 0 methods, this can hold values of any type, since every type implements 0 functions.
	var i interface{}
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	// Type Assertions provides access to an interface values's underlying concerete value
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float32)
	fmt.Println(f, ok)
	// Type Switch:
	// The type switch is a construct that permits several type assertions in the same switch statement.
	typeSwitch := func(i interface{}){
		switch v := i.(type){
		case int:
			fmt.Printf("%v * 2 == %v\n", v, v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		default:
			fmt.Println("Not sure what type is this is")
	}
}
	typeSwitch(100)
	typeSwitch("hello")
	typeSwitch(p)
	fmt.Println("===Methods on Types===")

	d := Developer{name: "Garvin"}
	doGreeting(d)
	fmt.Println(d.returnName())
	d.changeName("Madeline")
	fmt.Println(d.returnName())
	fmt.Println("===Error Handling===")

	/*
		// error type is a built-in interface type in the Go programming language.
		//	type error interface {
		//		Error() string
		//	}
	*/

	// Go is a run as happy path language.
	
	newDev := Developer{name: "Suraj"}
	err := newDev.setFavoritelanguage("C#")
	if err != nil{
		panic(err)
	}
	fmt.Println(newDev.greet())

}
