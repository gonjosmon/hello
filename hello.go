package main

import (
	"example/user/greetings"
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	message := greetings.Hello("Gonza")
	fmt.Println(message)
}
