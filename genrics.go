package main

//trying to implement genrics using interface
type Type interface{
	type 
		int,string

}
func Print[T Type](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}