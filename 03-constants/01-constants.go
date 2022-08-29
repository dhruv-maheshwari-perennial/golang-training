package main

import "fmt"

func main() {
	const a string = "10"
	const b = 10
	const (
		c = "String"
		d = 2.43
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
