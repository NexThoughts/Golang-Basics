package main

import "fmt"

func main() {
	emails:= make(map[string]string)

	emails["Office"] = "vijay@nexthoughts.com"
	emails["Personal"] = "vshukla684@gmail.com"
	fmt.Println(emails)

	delete(emails,"Personal")
	fmt.Println(emails)
}
