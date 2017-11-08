package main

import (
	"fmt"
	"log"
)

func main() {
	capitals := make(map[string]string) // If not dynamic can't add new entries
	capitals["Japan"] = "Tokyo"
	capitals["France"] = "Paris"

	v, ok := capitals["Narnia"]
	if !ok {
		// v is nil in this case
		log.Fatalf("No capital found for Narnia")
	}
	fmt.Printf("Narnia's capital is %s\n", v)
}
