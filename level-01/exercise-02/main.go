package main

import "fmt"

var (
	x int     // 0
	y string  // ""
	z bool    // false
)

func main() {
	fmt.Println(x, y, z)
	// A "Zero Value" is the name given to a value that the compiler assigns to a variable when a value isn't specified
}
