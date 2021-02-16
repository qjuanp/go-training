package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jp := person{
		firstName: "Juan",    // if attribute is no provided, go assumes based on the attribute order
		lastName:  "Giraldo", // order doesn't matter if attribute name is provided
	}

	fmt.Println(jp)
}
