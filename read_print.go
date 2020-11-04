package main

import(
	"fmt"
)

func main() {
	age := 37
	fmt.Printf("My age is %d\n", age)
	
	var edad2 int
	fmt.Println("Your age is: ")
	fmt.Scanf("%d\n", &edad2)
	fmt.Printf("Your age is %d\n", edad2)
}