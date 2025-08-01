package main

import (
	"fmt"
)

//	func fun( a, b int) int {
//		return a + b
//	}
type Books struct {
	tittle    string
	author_id int
	pages     int
	name      string
}

func main() {

	// var a int
	// var b int
	// fmt.Scanf("%d", &a)

	// fmt.Scanf("%d", &b)

	// ans := fun(a, b)
	// fmt.Println(ans)
	//   days := " Helloworld"
	// for index , value := range days {
	// 	fmt.Printf("val of index is %d , val of value is %c\n", index , value)

	// }
	var book1 Books
	var book2 Books

	book1.tittle = "Go Programming"
	book1.author_id = 1
	book1.pages = 1
	book1.name = "lesson of life"

	book2.tittle = "Go Programming"
	book2.author_id = 1
	book2.pages = 1
	book2.name = "lesson of life"

	fmt.Printf("book1 tittle is %s\n", book1.tittle)
	fmt.Printf("book1 author_id is %d\n", book1.author_id)
	fmt.Printf("book1 pages is %d\n", book1.pages)
	fmt.Printf("book1 name is %s\n", book1.name)

	fmt.Printf("book2 tittle is %s\n", book2.tittle)
	fmt.Printf("book2 author_id is %d\n", book2.author_id)
	fmt.Printf("book2 pages is %d\n", book2.pages)
	fmt.Printf("book2 name is %s\n", book2.name)

}
