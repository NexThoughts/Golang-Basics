package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Printf("Outer Sum is %d\n", sum)
	return func(num int) int {
		fmt.Printf("Before Sum is %d\n", sum)
		sum += num
		fmt.Printf("After Sum is %d\n", sum)
		return sum
	}
}

func main() {
	sum := adder()

	for i := 1; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
