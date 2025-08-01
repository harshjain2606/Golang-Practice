package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// var username string ="Harsh Jain"
	// number := 10
	// fmt.Printf("number  : %T\n", number)
	// fmt.Println("Welcome to our pizza app",number)

	// fmt.Println("Hello from", username)

	// Convert string to integer and add in int
	//  name := "123"
	//  sum := 12

	//  str,_ := strconv.Atoi(name)

	//   var res int = sum+ str
	//  fmt.Println(res)
	//  data := "1212"
	// parts := strings.Split(data, "9")
	// fmt.Println(parts)

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error while creating", err)
		return
	}

	defer file.Close()
	content := "Hello from Harsh Jain"
	_, error := io.WriteString(file, content)

	if error != nil {
		fmt.Println("error haii", error)
	}
}
