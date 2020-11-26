package main

import "fmt"

func main() {
	fmt.Print("Введите количество метров: ")
	var input float32
	fmt.Scanf("%f", &input)

	output := input / 0.3048

	fmt.Println("Количество футов: ", output)
}
