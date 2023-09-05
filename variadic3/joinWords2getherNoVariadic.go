package main

import "fmt"

func main() {

	// MAIN ADVANTAGE OF USING VARIADIC IS:
	// TO MAKE CODE MORE READABLE
	// A GO PROGRAM THAT JOINS WORDS TOGETHER WITH A SPECIFIED DELIMITER

	// 1. LET GAD MAKE IT WITHOUT USING VARIADIC FUNCTION
	var XLine string
	XLine = joinWords(",", []string{"Sam", "Gad", "Dan", "Fab"})
	fmt.Println(XLine)

	XLine = joinWords("/", []string{"Rwanda", "Kigali"})
	fmt.Println(XLine)

	XLine = joinWords("-", []string{"the", "first", "movie", "description"})
	fmt.Println(XLine)
}

func joinWords(delimiter string, theValues []string) string {
	var line string
	for index, val := range theValues {
		line = line + val
		// This -1 is to remove the last Delimiter like Gad,Dan,
		if index != len(theValues)-1 {
			line += delimiter // or line = line + delimiter
		}
	}
	return line
}
