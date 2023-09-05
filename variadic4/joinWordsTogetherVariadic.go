package main

import "fmt"

func main() {
	var line string

	// The change is here in passing...
	line = joinWords("-", "Rugby", "Is", "Harder", "Than", "Tennis")
	fmt.Println(line)

	line = joinWords(",", "One", "Two", "Three", "Four")
	fmt.Println(line)

	line = joinWords(" ", "Danny", "Trejo")
	fmt.Println(line)

	// EXPLODING ARGUMENTS

}

// G-Note: When defining any variadic function,
// Only the last parameter can be variadic.
func joinWords(delimiter string, values ...string) string {
	var line string
	for i, v := range values {
		line += v // same as: line = line + v
		if i != len(values)-1 {
			line += delimiter // same as: line = line + del
		}
	}
	return line
}
