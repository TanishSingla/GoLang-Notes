package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// fmt.Println("File Handling...");

	// file,_ := os.Create("example.txt");

	// // whenever we created a file we have to free its resources as well.
	// // This ensures that the file is closed properly after all other operations in the function have completed. Closing a file is important
	// //  because it releases the resources associated with it and ensures that any buffered data is written to the file.
	 
	// defer file.Close();

	// //adding content in a file...
	// content := "Hello there wassup "
	// byte,err := io.WriteString(file,content);
	// if err!=nil{
	// 	fmt.Println("Content added successfully")
	// }else {
	// 	fmt.Println("Added data is of",byte,"Bytes");
	// }


	// reading file data :- 
	
	file,err := os.Open("example.txt");
	if err!=nil{
		fmt.Println("Error while opening the file..",err);
		return
	}
	defer file.Close()

	//creating buffer to read file content
	buffer := make([]byte,1024)

	//read the file content into buffer
	for{
		n,err2 := file.Read(buffer);
		if err2 == io.EOF{
			break
		}
		if err2!=nil {
			fmt.Println("error while reading data",err2);
			return
		}
		fmt.Println(string(buffer[:n]));
	}


	//easy way for reading a file
	//its not recommended to use direct Readfile , as for very large
	//file it loads whole file in memory might get memory issue
	//so its better to use buffer
	//it will load data chunks by chunks
	file2,err3 := os.ReadFile("example.txt");
	if err3!=nil {
		fmt.Println(err3)
	}else{
		fmt.Println(file2)
		fmt.Println(string(file2))
	}
}