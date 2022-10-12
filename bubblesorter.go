package main

import (
	"fmt"
)

func main() {
	var numbers [256]int

	var numbersToSort = 0
	var temp int = 0

	fmt.Print("How many numbers to sort: ")
	fmt.Scanln(&numbersToSort)

	for i := 0; i < numbersToSort; i++ {
		fmt.Printf("Num #%d: ", i + 1)
		fmt.Scanln(&numbers[i])
	}

	for x := 0; x < numbersToSort; x++ {
		for y := 0; y < numbersToSort - x - 1; y++ {
			if numbers[y] > numbers[y + 1] {
				temp = numbers[y]
				numbers[y] = numbers[y + 1]
				numbers[y + 1] = temp
			}
		}
	}

	fmt.Println("\nAfter sorting:")
	for i := 0; i < numbersToSort; i++ {
		fmt.Printf("%d ", numbers[i])
	}
}