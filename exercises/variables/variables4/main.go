// variables4
// Make me compile!

package main

import "fmt"

func main() {
	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
	x := "TEN" // Don't change this line
	fmt.Printf("x has the value %s", x)

	if true {
		x = "1"
		fmt.Println(x + "1")
	}

	fmt.Println(x)
}
