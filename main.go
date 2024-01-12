package main

import "fmt"

func main() {
	var name string = "John"
	var name2 *string = &name // => pointer to name
	*name2 = "Jean"           // => change value of name in memory address
	// or name := "John"
	fmt.Printf("Hello %#v : %#v\n", name, *name2)

	// Array and slice
	a := [3]int{1, 2, 3}
	b := a
	b[1] = 4
	c := a[1:] // Display with first element
	fmt.Printf("Array: %#v : %#v\n", a, b)
	fmt.Printf("Array first element: %#v : %#v\n", a, c)

	d := []int{1, 2, 3}                    // slice
	fmt.Printf("Array len: %#v\n", len(d)) // length of slice
	fmt.Printf("Array cap: %#v\n", cap(d)) // capacity of the slice

	e := append(d, 4) // push new element in the slice
	fmt.Printf("Array push: %#v : %#v\n", d, e)

	// Objects and maps
	maps := map[string]string{
		"title":       "Faire le tutoriel",
		"description": "Parler le golang",
	}
	fmt.Printf("Map: %#v : %#v\n", maps, maps["title"])
	maps["title"] = "Paul"
	fmt.Printf("Map: %#v\n", maps)

}
