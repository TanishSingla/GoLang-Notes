package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Data Conversion..");

	// we can only perform calculations or any operations on same type of
	// datatype.

	var a int  = 10;
	var b float64 = 21.3;

	// we can not add a int to a flaot
	// either we have to convert an int to float
	// or a float to int.
	fmt.Println(float64(a)+b);

	// converting a int to string
	var c int = 1231;
	fmt.Printf("Type of variable is %T \n",c);
	
	// Integer to string.
	var s string  = strconv.Itoa(c);
	fmt.Printf("Type of variable is %T \n",s);

	// String to Integer.
	s2 := "123"
	num_to_str,_ := strconv.Atoi(s2);
	fmt.Printf("Type -> %T",num_to_str);
}