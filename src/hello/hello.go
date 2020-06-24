package main

import (
	"example.com/user/hello/morestrings"
	"example.com/user/hello/test"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	test.Test()
	Test()
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
