package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	age int
}

func main() {
	fmt.Println("STRUCTS in Go...");

	var A Person;
	A.FirstName = "Xavier"
	A.LastName = "Adams"
	A.age = 32
	fmt.Println("Person -> A",A);
	
	A.FirstName = "Changed"
	fmt.Println("Person -> A",A);
	
	B := Person{
		FirstName: "new person b",
		LastName: "lastname",
	}
	fmt.Println("Person -> B",B);

	var C = new(Person)
	C.FirstName = "Person C"
	fmt.Println(C);
}