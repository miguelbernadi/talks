package main

import "fmt"

var me string

func main() {
	me = "me"
	you := me + " and you"
	fmt.Println("Hello World!", you)
}
