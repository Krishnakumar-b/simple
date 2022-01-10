package main

import (
	"fmt"
)

func calculate(n int) int {

	var sum, rem int
	for n != 0 {
		rem = n % 10
		sum = sum + rem
		n = n / 10
	}

	return sum
}

func main() {
	n := 9999999999999
	number := calculate(n)
	for {
		if number > 9 {
			number = calculate(number)
		}
		break
	}
	fmt.Println(n, number)
}
