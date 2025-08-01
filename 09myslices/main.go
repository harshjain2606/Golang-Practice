package main

import (
	
	"fmt"
	
)
// func greet(){
// 		fmt.Println("Hello, welcome to the Go programming world!")
// 	}c

// 	func add(a int , b int )  int  {   //  signature of the function
// 		// This function takes two integers and returns their sum
// 		return a+b
// 	}

func main()  {

// highScore := make([]string, 4)
// highScore[0] = "10"
// highScore[1] = "20"
// highScore[2] = "30"
// highScore[3] = "40"
// highScore = append(highScore," 50")
// fmt.Println("High Scores:", highScore)
// 		fmt.Printf("Type of highScore: %T\n", highScore)

// User input example
		/* reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your name: ")
		input, _ := reader.ReadString('\n')
		fmt.Println("Your name is:", input)

*/

		type User struct {
			Name     string
			Age      int
			Email    string
			IsActive bool

		}
	/*	hitesh := User{"Harsh Jain", 23,"jaainharsh383@gmail.com", true}

		fmt.Println("User Details:", hitesh)
		fmt.Printf("Type of User: %+v\n", hitesh)


	if num := 10 
		num < 20 {
			fmt.Println("Number is less than 20")
		} else {
			fmt.Println("Number is 20 or more")
		}
			
		*/

	//  days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	//  fmt.Println("Days of the week:",days)

	//  for d := 0; d<len(days); d++{
	// 	 fmt.Println("Day", d+1, ":", days[d])
	//  }


	// for i := range days{
	// 	fmt.Println("Day", i+1, ":", days[i])
	// }


// 	for i, day := range days {
// 		fmt.Printf("i is %v and day is %v\n", i, day)
// }

//  greet()

//  result := add(5, 10)
	
// fmt.Println("Sum of 5 and 10 is:", result)

// reader := bufio.NewReader(os.Stdin)
//  	fmt.Println("Enter your name: ")
// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Your name is:", input)
   var  age int ;
// 	fmt.Println("Enter your age: ")
fmt.Scanf("%d", &age)
    var tall int = age
	fmt.Println("Your age is:", tall)
 	// if age>45 {
	// 	fmt.Println("You are too old to play this game",input)

	// } else{
	// 	fmt.Println("You are young enough to play this game")
	// }


// defer fmt.Println("Welcome to our pizza app"). // defer follows last in first out principle
//   fmt.Println("This is my slices app")


	

}