package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter Your Name :- " );
	// var Name string;
	// fmt.Scan(&Name);
	// println("Hello",Name)
	
	
	reader := bufio.NewReader(os.Stdin)
	Name,_ := reader.ReadString('\n');
	println("Hello",Name)
}