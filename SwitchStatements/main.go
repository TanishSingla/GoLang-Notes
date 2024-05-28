package main

import "fmt"

func main() {

	fmt.Println("Enter a number");
	var day int;
	fmt.Scanf("%d",&day);
	var s string
	switch day{
	case 1:
		s = "Monday"
	case 2:
		s = "Tuesday"
	case 3:
		s = "Wednesday"
	case 4:
		s = "Thursday"
	case 5:
		s = "Friday"
	case 6:
		s = "Saturday"
	case 7:
		s = "Sunday"
	case 8,9,10:
		s = "Multple conditions in single case"
	default:
		s = "Enter the wrong number 😐"
	}
	fmt.Println("--->",s);

}