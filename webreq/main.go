package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web requests...")
	resp,err := http.Get("https://jsonplaceholder.typicode.com/todos/1");
	// type of resp is of response type , http.response
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close();
	// fmt.Println(resp)
	//read resp

	data,err2 := ioutil.ReadAll(resp.Body);
	if err2!=nil{
		fmt.Println("Error while reading response",err)
	}
	fmt.Println(string(data))


}