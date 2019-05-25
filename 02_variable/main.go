package main

import "fmt"

func main() {
	//MAIN Types
	// string
	// bool
	//int
	//int8, int16, int32, int64
	//uint, uint8, uint16, uint32, uint64, uintptr
	//byte -alias for uint8
	//rune- alias for uint32
	//float32 float64
	//complex64, complex128

	//Using var
	var name string = "Vijay Shukla"
	fmt.Println(name)

	//Using var
	var name1 = "Vijay Shukla"
	fmt.Println(name1)

	name2 := "Vijay Shukla"
	fmt.Println(name2)

	var age int = 31
	fmt.Println(age)

	fmt.Println(name, age)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)

	name4, email := "Vijay Shukla", "vijay@nexthoughts.com"

	fmt.Println(name4)
	fmt.Println(email)
}
