package main

import (
	
	"fmt"
	"io"
	"net/http"
)

func main()  {
	fmt.Println("welcome CrUD")

res, err :=	http.Get("https://jsonplaceholder.typicode.com/todos/1")
if(err!=nil){
	fmt.Println("error present",err)
	return
}
  defer res.Body.Close()
 data , err:=io.ReadAll(res.Body)
 if(err!=nil){
	fmt.Println("error is present",err)
	return
 }
 fmt.Println( string(data))
   

 

   
}
