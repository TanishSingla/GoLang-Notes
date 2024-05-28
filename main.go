package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World");

	// defer keyword...
	//defer is always in starting of a function
	//and that particular function will run at the end of the main function
	//means at the end of the program
	//and if we are using defer more then once
	//then then get called in LIFO order (means they are getting stored in stack)
	fmt.Println("Firsttt"); 
	defer fmt.Println("Second");
	defer fmt.Println("Third");
	fmt.Println("Fourth");

}