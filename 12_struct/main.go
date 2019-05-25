package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	/*firstName string
	lastName  string
	city      string
	gender    string
	age       int*/

	firstName, lastName, city, gender string
	age                               int
}

//Value Receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + ". my age is " + strconv.Itoa(p.age)
}

//Pointer Receiver
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//person1 := Person{firstName: "Vijay", lastName: "Shukla", city: "Noida", gender: "f", age: 31}
	person1 := Person{"Vijay", "Shukla", "Noida", "f", 31}

	fmt.Println(person1)
	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Dikshit")
	fmt.Println(person1.greet())
}
