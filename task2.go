package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	count := len(x)
	element := x[count-1]
	for count > 0 {
		count--
		minimum := x[count]
		if element > minimum {
			element = minimum
		}
	}
	fmt.Println("Минимальное число: ", element)
}
