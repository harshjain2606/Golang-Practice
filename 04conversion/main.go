package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

type Contact struct {
	Email string
	Phone string
}

type  Address struct {
	House int
	Area  string
	State  string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address  Address
}


func main()  {
	 var employee1 Employee
	 employee1.Person_Details =Person{
		FirstName:  "Harsh Jain",
		LastName:  "Jain",
		Age:  12,
	 }
   employee1.Person_Contact.Email = "jaainharsh383@gmailcom"
   employee1.Person_Contact.Phone = "4343424242"

   employee1.Person_Address = Address{
	  House : 1,
	  Area : "123",
	  State : "delhi",
   }

   fmt.Println("employee1 contact",employee1.Person_Contact)
    fmt.Println(" emolyee details", employee1.Person_Details)
	fmt.Println("employee adress", employee1.Person_Address)
	
}