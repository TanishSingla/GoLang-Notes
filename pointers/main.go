package main

import "fmt"

func modifyByReference(x*int){
	*x = *x + 10;
}

type struct1 struct{
	x int;
	y int;
}

func main() {

	fmt.Println("Pointers in Go..")

	 a  := 10;
	 b   := &a
	fmt.Println(b);

	var ptr *int = &a;
	fmt.Println(*ptr)

	// pass by reference
	fmt.Println(a)
	modifyByReference(&a)
	fmt.Println(a)

	str1 := struct1{1,2};
	ptrForStrct := &str1
	// Note :- (*ptrForStrct).x is same as ptrForStrct.x
	(*ptrForStrct).y  = 9
	ptrForStrct.x = 19
	fmt.Println(str1);

}