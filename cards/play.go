package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	numbers1 := append(numbers, 7, 8, 9)

	for num := range numbers {
		fmt.Println("Number Slice", num*2)
	}

	for num := range numbers1 {
		fmt.Println("Numbers 1 Slice", num+1)
	}

}
