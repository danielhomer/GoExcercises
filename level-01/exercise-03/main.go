package main

import "fmt"

var (
	x int = 43
	y string = "James Bond"
	z bool = true
)

func main() {
	s := fmt.Sprintf("%v%v%v", x, y, z)
	fmt.Println(s)
}
