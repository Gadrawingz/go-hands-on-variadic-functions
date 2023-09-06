package main

import "fmt"

// Reference: https://golangbot.com/variadic-functions/
func main() {

	// When variadic function has been used...
	findNumber(49)
	findNumber(700, 34, 55, 63, 70, 65, 100, 121)
	findNumber(100, 34, 55, 63, 70, 65, 100, 121)

	// When Slice was used instead of variadic function
	findNumberUsingSlice(8, []int{})
	findNumberUsingSlice(51, []int{34, 40, 44, 49, 50, 51, 55, 63, 99})
}

func findNumber(number int, numbers ...int) {
	fmt.Println("***********************************")
	fmt.Printf("Type of numbers is %T\n", numbers)
	found := false
	for index, value := range numbers {
		if value == number {
			fmt.Println("Number", number, " is found at index ", index, "in", numbers)
			found = true
		}
	}

	if !found {
		fmt.Println("Number", number, "is not found in", numbers)
	}
	fmt.Printf("\n")
}

func findNumberUsingSlice(number int, numbers []int) {
	fmt.Println("--------------------------------")
	fmt.Printf("Type of numbers is %T\n", numbers)
	found := false
	for index, value := range numbers {
		if value == number {
			fmt.Println("Number", number, " is found at index", index, "in", numbers)
			found = true
		}
	}

	if !found {
		fmt.Println("Number", number, "is not found in", numbers)
	}
	fmt.Printf("\n")

}
