package main

import (
	"fmt"
)

func calculate(n int64) int64 {

	var sum, rem int64

	for n != 0 {
		rem = n % 10
		sum = sum + rem
		n = n / 10
	}

	return sum
}

func main() {
	var n int64

	var arrNum [8]int64

	fmt.Println("enter a number of maximum length 8")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Print(err)
	}
	i := 0

	num := n
	for num != 0 {
		arrNum[i] = num % 10
		num = num / 10
		i++
	}

	inttostring := fmt.Sprintf("%s", arrNum)
	fmt.Printf("%T\n", inttostring)

	number := calculate(n)
	for {
		if number > 9 {
			number = calculate(number)
		}
		break
	}
	fmt.Println(number)

}
