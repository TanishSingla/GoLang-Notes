package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int32   `json:"ageOfPerson"`
	Isadult bool 

}

func main() {
	fmt.Println("How to convert data into json format...")

	p1 := Person{
		Name: "First",
		Age: 21,
		Isadult: true,
	}

	//convert person data into json format (marshalling)
	jsonData,err := json.Marshal(p1);
	fmt.Printf("Type of converted data is %T \n",jsonData);

	if err!=nil{
		fmt.Println("Error while converting data to JSON")
		return
	}
	// fmt.Println(string(jsonData))
	// fmt.Printf("Type of converted data %T \n",string(jsonData))

	//decoding unmarshalling...
	var personData Person;
	err2 := json.Unmarshal(jsonData,&personData);

	if err2==nil{
		fmt.Println(personData);
	}

}