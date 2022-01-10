package main

//trying to implement genrics using interface
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}