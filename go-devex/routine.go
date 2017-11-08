package main

import (
	"fmt"
)

func main() {
	routine := func() {
		fmt.Println("I'm a goroutine!")
	}
	go routine()
	fmt.Println("Now I'm done")
}
