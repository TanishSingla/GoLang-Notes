package main

import "fmt"

func main() {
	fmt.Println("In Go language there is only for loop");

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i);
	// }

	//Infinite loop
	// for{
	// 	fmt.Println("Infinite Loop");
	// }


	//range keyword
	nums := []int{1,2,3,4,5,6};
	for i,value := range nums {
		fmt.Printf("Index is %d and value is %d \n",i,value);
	}

	s := "Learning Go Programming...";
	for c,ch := range s{
		fmt.Printf("Index is %d and character is %c \n",c,ch);
	}
}