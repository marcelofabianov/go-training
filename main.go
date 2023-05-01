package main

import "fmt"

func main() {
	print("\nHello ", "World ", 2022, " Go!")
	println("Hello", "World ", 2022, " Go!")

	///////////////////lib fmt////////////////////////
	fmt.Println("Hello", "World ", 2022, " Go!")

	s := fmt.Sprintf("%s", "string-here")
	fmt.Printf("%s%s", s, "\n")

	f := fmt.Sprintf("%0.2f", 34.56756)
	fmt.Println("float:", f)
}
