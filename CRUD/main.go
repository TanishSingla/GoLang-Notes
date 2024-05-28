package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type task struct {
	UserId    int    `json:"userId`
	ID        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

func GetReq() {
	fmt.Println("------------------------GET-------------------------")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error while fetching data :- ", resp.StatusCode)
	}
	// data,err := ioutil.ReadAll(resp.Body);
	// if err!=nil{
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Print(string(data))

	var Todo1 task
	err = json.NewDecoder(resp.Body).Decode(&Todo1)

	if err != nil {
		fmt.Println("error while decoding data")
		return
	}
	fmt.Println(Todo1)
	fmt.Println(Todo1.ID)
	fmt.Println(Todo1.Title)
}

func PostReq() {
	fmt.Println("------------------------POST-------------------------")

	dataToAdd := task{
		UserId:    99,
		Title:     "Data added",
		ID:        1,
		Completed: false,
	}

	//converting data to json
	jsonData, err := json.Marshal(dataToAdd)
	if err != nil {
		fmt.Println("Error while converting data to json", err)
		return
	}

	//convert json data to string...
	jsonStringData := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonStringData)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Post(myUrl, "applications/json", jsonReader)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response data", string(data))
}

func PutReq() {
	fmt.Println("------------------------PUT-------------------------")
	dataToUpdate := task{
		UserId:    1,
		Title:     "Data Updated",
		ID:        1,
		Completed: false,
	}

	jsonData, err := json.Marshal(dataToUpdate)
	if err != nil {
		fmt.Println("Error while converting data to JSON", err)
		return
	}

	jsonReader := strings.NewReader(string(jsonData))

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	req.Header.Set("Content-type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	//send req
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func DelReq() {
	fmt.Println("------------------------DELETE-------------------------")

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	//send req
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error while making request", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}

func main() {
	fmt.Println("...CRUD Operations...")
	GetReq()
	PostReq()
	PutReq()
	DelReq()
}
