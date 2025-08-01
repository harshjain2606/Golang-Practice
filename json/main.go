package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	Userid    int    `json:"name"`
	Id        int    `json:"id"`
	Tittle    string `"json:tittle"`
	Completed bool   `"json:Completed"`
}

// func PerformGetRequest() {
// 	fmt.Println("we are learning about Json")
// 	person := Person{Name: "Harsh Jain", Age: 12, IsAdult: true}
// 	fmt.Println("ans is", person)

// 	jsonData, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("error in data", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))

// 	//unmarshaling

//		var persondata Person
//		err = json.Unmarshal(jsonData, &persondata)
//		if err != nil {
//			fmt.Println(" ", err)
//			return
//		}
//		// fmt.Println(" ",persondata)
//	}
func performPostRequest() {
	todo := Todo{
		Userid:    23,
		Tittle:    "Harsh Jain",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error in", err)
		return
	}

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)
	myurl := "https://jsonplaceholder.typicode.com/todos"
	// send post req
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}

func main() {
	performPostRequest()

}
