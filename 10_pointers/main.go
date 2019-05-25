package main

import (
	"fmt"
)

func main() {
	a:=5
	b:=&a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
	fmt.Println(*&a)
	fmt.Println(*&b)

	*b = 10
	fmt.Println(a)
	fmt.Println(b)
}
