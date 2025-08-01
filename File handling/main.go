package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

	// content := "Hello from Harsh Jain"
	// _, error := io.WriteString(file, content)

	// if error != nil {
	// 	fmt.Println("error haii", error)
	// }

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error in opening file", err)
		return
	}
	defer file.Close()
	// create buffer to read the file content
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error", err)
			return
		}

		// proccess  the read content
		fmt.Println(string(buffer[:n]))
	}

	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("error in reading file", err)
		return
	}
	fmt.Println("", content)
	fmt.Printf("%T", content)
}
