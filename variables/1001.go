package main

import (
	"fmt"
)

// method 4:
var r int = 99

func main() {
	// intialize a variable
	// method 1:
	var i int
	i = 42
	fmt.Println(i)

	// method 2
	var j int = 23
	fmt.Println(j)

	// method 3
	k := 23
	fmt.Println(k)
	// here `:=` makes golang decide the variable type on it's own

	// methode 4 (outsite main function, can't use method 3 at that place)
	fmt.Println(r)

	// some limitations of methode 3 is
	// can't make float32 type variable, so for more control use methode 2
	m := 9
	fmt.Printf("%v, %T \n", m, m)
	n := 9.
	fmt.Printf("%v, %T \n", n, n)

	var t float32 = 24
	fmt.Printf("%v, %T \n", t, t)
}
