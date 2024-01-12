package main

import "fmt"

func main() {
	var name string = "John"
	var name2 *string = &name // => pointer to name
	*name2 = "Jean"           // => change value of name in memory address
	// or name := "John"
	fmt.Printf("Hello %#v : %#v", name, *name2)
}
