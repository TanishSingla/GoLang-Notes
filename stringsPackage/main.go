package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Strings package in Go....")
	
	fmt.Println("---------------Split----------");
	data := "abc,def,ghi"
	parts := strings.Split(data, ",")
	fmt.Println(parts)
	fmt.Printf("Type of parts is %T \n",parts)

	fmt.Println("---------------Count----------");
	data = "12312312112"
	fmt.Println(strings.Count(data,"1"));
	fmt.Println(strings.Count(data,"9"));

	fmt.Println("---------------TrimSpace----------");
	data = "            hello     boom       ";
	data = strings.TrimSpace(data)
	fmt.Println(data);
	
	fmt.Println("---------------Join----------");
	fmt.Println(strings.Join([]string{data,"joined","strings"}," "));

	fmt.Println("joining " + " strings " + " using + op ");

}