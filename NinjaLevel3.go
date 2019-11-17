// Exercises - Ninja Level 3

// hands 1

package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}


// hands 2

package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {

		fmt.Println(i)
		for j := 0; j < 3; j++ {

			fmt.Printf("\t %#U \n", i)
		}
	}

}
//hands 3

package main

import (
	"fmt"
)

func main() {

	bd := 1990
	for bd < 2019 {
		fmt.Println(bd)
		bd++
	}
}

//hands 4

package main

import (
	"fmt"
)

func main() {

	bd := 1990
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

// hands 5

package main

import (
	"fmt"
)

func main() {
	num := 10
	for {
		if num >= 101 {
			break
		}

		fmt.Println(num % 4)
		num++

	}
}

//hand 6 and 7 

package main

import (
	"fmt"
)

func main() {
	num := 10
	if num == 11 {
		fmt.Println("equal to 10")
	} else if num == 12 {
		fmt.Println("not equal to 12 ")
	} else {
		fmt.Println("equal  to ", num)
	}

}
// hand 8

package main

import (
	"fmt"
)

func main() {
	switch {

	case true:
		fmt.Println("it's true")
	case false:
		fmt.Println("it's false")
	}

}
// hands 9 

package main

import (
	"fmt"
)

func main() {
	favSport := "Surfing"
	switch favSport {

	case "running":
		fmt.Println("running")
	case "boxing":
		fmt.Println("boxing")
	case "paragliding":
		fmt.Println("paragliding")
	case "Surfing1":
		fmt.Println("Surfing")
	default:
		fmt.Println("Surfing")
	}

}

