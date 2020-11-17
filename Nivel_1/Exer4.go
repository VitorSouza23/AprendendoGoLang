package main

import "fmt"

type tardis int

var x tardis

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%v", x)
}
