package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL FUNCTIONALITY....")

	myurl := "https://api.publicapis.org/entries?key=1"
	fmt.Printf("Type of myurl is %T \n",myurl);

	convertedUrl,err := url.Parse(myurl);
	if err!=nil{
		fmt.Println("Error while parsing");
		return;
	}
	fmt.Printf("Type of convertedUrl using parse is %T \n",convertedUrl);
	fmt.Println(convertedUrl.Scheme) //http or https
	fmt.Println(convertedUrl.Host)
	fmt.Println(convertedUrl.Path) 
	fmt.Println(convertedUrl.RawQuery) // query parameters

	convertedUrl.Host = "abcdef"
	convertedUrl.Scheme = "http"

	changedUrl := convertedUrl.String()
	fmt.Println(changedUrl)


	

	
}