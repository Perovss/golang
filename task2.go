package main

import "fmt"

var y int

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	count := len(x)
	for count > 1 {
		count = count - 1
		y := x[count]
		fmt.Println(y)
	}

}
