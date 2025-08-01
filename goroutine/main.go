package main

import (
	"fmt"
	"time"

)
  func hi(){
	fmt.Println("hi")
  }
 func hello(){
	time.Sleep(900*time.Millisecond)
	 fmt.Println("Hey Harsh ")
	 
	
 }
func main()  {

	 go hi();
	go hello();
	time.Sleep(1000*time.Millisecond)
	
}