package main

import "fmt"

func main() {
	integers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range integers {
		if n%2 == 0 {
			fmt.Println(fmt.Sprintf("%d is even", n))
		} else {
			fmt.Println(fmt.Sprintf("%d is odd", n))
		}
	}
}
