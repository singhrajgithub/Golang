
package main

import (
	"fmt"
)

func main() {
	g := 42 == 42
	h := 42 <= 42
	i := 42 >= 42
	j := 42 != 42
	k := 42 < 42
	l := 42 > 42

	fmt.Println(g, h, i, j, k, l)

}
