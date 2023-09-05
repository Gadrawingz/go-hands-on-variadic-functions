package variadic1

import "fmt"

func main() {

	// Uses of Our Variadic function
	sendGreetings("Dave")
	sendGreetings("Sammy", "Diana", "Kate")
	sendGreetings("Moise", "Wayne")

}

// Ref: https://www.digitalocean.com/community/tutorials/how-to-use-variadic-functions-in-go
func sendGreetings(names ...string) {
	for _, name := range names {
		fmt.Printf("Wassup %s!\n", name)
	}
	fmt.Println("***********************")
}
