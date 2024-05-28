package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Date And Time","02-01-2006");

	TodaysDate := time.Now();
	fmt.Println(TodaysDate); 
	formattedDate := TodaysDate.Format("02-01-2006,Monday");
	fmt.Println(formattedDate);
	formattedDate = TodaysDate.Format("02-01-2006,15:04:05");
	fmt.Println(formattedDate);
	formattedDate = TodaysDate.Format("2006/01/02,15:04:05");
	fmt.Println(formattedDate);
	formattedDate = TodaysDate.Format("2006/01/02,3:04 PM");
	fmt.Println(formattedDate);

	//Converting a string to date format
	str := "2024-05-28"
	convertedStr,_ := time.Parse("2006-01-02",str);
	fmt.Println(convertedStr);
	 

}