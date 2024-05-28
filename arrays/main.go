package main

import "fmt"

func main() {

	
	var arr1 [5]string
	arr1[0] = "abv"
	arr1[4] = "dsfds"
	fmt.Println(arr1);
	fmt.Println(arr1[1]);
	fmt.Printf("%q\n",arr1);

	arr2 := [5]int{1,2,3,4,5};
	fmt.Println(arr2);

	var arr3[5] bool;
	arr3[1] = true;
	fmt.Println(arr3);

	var arr4[5] int;
	fmt.Println("arr4 ->",arr4);
	fmt.Println("Length of arr3 is :",len(arr3))
}