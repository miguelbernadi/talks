package main

import "fmt"

func main() {
	// START OMIT
	func(s string) {
		fmt.Println(s)
	}("Hi!")
	// END OMIT
}
