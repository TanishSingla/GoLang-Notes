package main

import "fmt"

func main() {
	fmt.Println("Learning maps..");


	m := make(map[string]int);
	m["first"] = 100;
	m["second"] = 98;
	m["third"] = 54;
	m["fourth"] = 44;

	fmt.Println(m["first"]);
	fmt.Println(m["second"]);

	for key,value := range m{
		fmt.Printf("Key -> %s <-----> value -> %d \n",key,value);
	}

	//to check if any key is present in the map or not...
	_,present := m["first"];
	fmt.Println(present);
	_,present = m["abc"];
	fmt.Println(present);

	val,pre := m["first"];
	fmt.Println(val,pre);

	//deleting keys from map
	//note if any key is not present in the map 
	//and we try to access its value then it will return 0
	delete(m,"first");
	println(m["first"],"---");

	// we can initialize the map while declaration as well..

	m2 := map[string]int{
		"Abc":3,
		"Defg":4,
	}

	fmt.Println(m2);
}