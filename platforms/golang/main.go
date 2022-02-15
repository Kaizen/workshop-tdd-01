package main

import "fmt"

func main() {
	fmt.Printf("foo = %d", foo())
}

func foo() int {
	return 4
}
