package main

import (
	"fmt"
	"github.com/vijayshukla30/go_crash_course/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))

	fmt.Println(strutil.Reverse("Vijay Shukla"))

}
