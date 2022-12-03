package main

import (
	"fmt"
)

// But while declaring as a Package level var, we've to declare as Syntax3
var I int = 22

//While declaring multiple variables, we can block it together.

var (
	name   string = "skv"
	age    int    = 19
	degree string = "BTech"
)

func main() {
	//fmt.Println(I)
	/*
		//Syntax1
		var i int
		i = 22

		//Syntax2
		i := 22
		//Syntax3
		//var i int = 22
	*/

	fmt.Println("Hello World!")

	var i int = 02 // This is what Shadowing, that is updating Package var to new val.

	// We can't declare same var multiple times in a Scope.
	//i := 23
	/*
		//j := 11
		We'll get error, when we not used the declared var. It's a great feature in Go.
	*/
	fmt.Println(i)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(degree)

}
