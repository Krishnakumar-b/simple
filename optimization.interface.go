package main

import (
	"fmt"
)

type DataType interface {
}

func calculat(T DataType) error {

	//fmt.Println(T)
	xtype := fmt.Sprintf("%T", T)
	//fmt.Println("*****", xtype)
	if xtype == "int" {

		n := T.(int)

		for n/10 != 0 {
			n = n%10 + n/10
		}
		fmt.Printf(" Type : %T\n", n)
		fmt.Println("Number : ", n)
		fmt.Println("Sum : ", n)

	}
	if xtype == "string" {
		fmt.Println("Error!Cannot use string values")

	}

	if xtype == "float64" {
		fmt.Println("Error!Cannot use float values")
	}

	return nil
}
func main() {
	var i DataType

	i = 9999
	calculat(i)

	i = "hello"
	calculat(i)

	i = 11111.10
	calculat(i)
}
