package variadic2

import "fmt"

func main() {

	// Using Function in action...
	sayMorningToPeople()
	sayMorningToPeople("Fabrice")
	sayMorningToPeople("La4ouine", "Sindy")

}

func sayMorningToPeople(names ...string) {
	if len(names) == 0 {
		fmt.Println("Okay,Nobody to greet now!")
		return
	}
	for _, name := range names {
		fmt.Printf("Morning %s !!\n", name)
	}
}
