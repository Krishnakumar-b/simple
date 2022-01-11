package main

import (
	"fmt"
	"os"
	"strconv"
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
	var n int
	var arrNum [8]int
	x := 015
	fmt.Printf("%T", x)
	fmt.Println("enter a number of maximum length 8")
	_, err := fmt.Scanf("%d", &n)

	if n > 99999999 {
		fmt.Println("Error")
		os.Exit(0)
	}
	if n%1 != 0 {
		fmt.Println("Error : Cant accept floating point number")
		os.Exit(0)
	}
	if n <= 0 {
		fmt.Println("Error : Cant accept negative number")
		os.Exit(0)
	}

	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("................\n%T value= %s\n", strconv.Itoa(int(n)), strconv.Itoa(int(n)))
	i := 0

	num := n
	for num != 0 {
		arrNum[i] = num % 10
		num = num / 10
		i++

	}

	number := calculate(x)
	for {
		if number > 9 {
			number = calculate(number)
		}
		break
	}
	fmt.Println(number)

}
