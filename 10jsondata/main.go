package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	UserId    int    `json:"userId"`
	Title     string `json:"title"`      
	Completed bool   `json:"completed"`  
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var users []User 

	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	for i, user := range users {
		
		fmt.Printf("id : %d, title: %s, completed: %t\n", i+1, user.Title, user.Completed)
	}
	file, err := os.Create("save.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	

	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(" JSON data saved to todos.json using Encoder")
}

