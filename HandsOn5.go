
package main

import (
	"fmt"
)

type myvar int

var x myvar
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 30
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T \n", y)
	test() // to check difference in the y global and y block value i created differnt fun below and called that from main func()

}

func test() {
	fmt.Println(y)
}
