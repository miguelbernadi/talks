package main

import "fmt"

var hello = func() string {
	return "Hello"
}

var goodbye = func() string {
	return "goodbye"
}

func Say(s func() string) {
	fmt.Println(s())
}

func main() {
	Say(hello)
	Say(goodbye)
}
