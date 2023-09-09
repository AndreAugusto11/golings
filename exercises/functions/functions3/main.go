// functions3
// Make me compile!

package main

import "fmt"

func main() {
	var x = 2
	call_me(x)
}

func call_me(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
