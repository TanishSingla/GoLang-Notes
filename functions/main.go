package main

import "fmt"

func Add(a int ,b int) int{
	return a+b
}

// lets say there are several parameters with same datatype , in that case we can
// just specify the datatype in the last parameter.

func Add2(a,b,c,d,e int) ( result int){
	result =  a+b+c+d+e
	return
}
func main() {

	fmt.Println("Learning Functions")
	fmt.Println(Add(2,3));
	fmt.Println(Add2(1,2,3,4,3));
}