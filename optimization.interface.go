// package main

// import (
// 	"fmt"
// 	"go/types"
// )

// type DataType interface {
// }

// func calculat(T DataType) error {

// 	fmt.Println(T)
// 	xtype := fmt.Sprintf("%T", T)
// 	fmt.Println("*****", xtype)

// 	if xtype == types.Int {
// 		fmt.Println("string")
// 	}
// 	return nil
// }
// func main() {
// 	var i interface{}

// 	i = 42
// 	calculat(i)

// 	// i = "hello"
// 	// calculat(i)
// }
