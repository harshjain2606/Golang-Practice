package main 
import "fmt"
func main()  {
	number := 23
	
	fmt.Printf("Value at pointer address %T \n", number) 
	var ptr *int = &number // Pointer to the variable 'number'
	fmt.Println("Value of number:", ptr)
	fmt.Println("Address of number:", &ptr)
	fmt.Printf("Value at pointer address %T", *ptr) // Dereferencing the pointer to get the
}
