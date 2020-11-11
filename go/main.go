package main

import(
	"fmt"
)

func hello() string {
	return "hello"
}

func main() {
	str := hello()
	fmt.Printf("%s\n", str)
}
