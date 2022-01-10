package main

import (
	"fmt"
)

func calculate(n int64) int64 {
	var arrNum [3]int
	var sum, rem int64

	i := 0

	for n != 0 {
		rem = n % 10
		if rem == 0 {
			fmt.Print("error")
		}
		arrNum[i] = int(rem)
		i++
		sum = sum + rem
		n = n / 10
	}
	// var mapNumber map[int]int
	// mapNumber=arrNum
	// for i=len(arrNum);i>=0;i-- {

	// }
	return sum
}

func main() {
	var n int64
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Print(err)
	}

	number := calculate(n)
	for {
		if number > 9 {
			number = calculate(number)
		}
		break
	}

	fmt.Println(n, number)

}
