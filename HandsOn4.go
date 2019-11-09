
package main

import (
	"fmt"
)

type myvar int

//var b int

var x myvar

func main() {

	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println("\n After Assigment", x)
	//	b = x
}
