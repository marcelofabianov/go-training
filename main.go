package main

import "fmt"

// variavel escopo de pacote
var name  string

func main() {
	name = "Marcelo fabiano"

	fmt.Println("Hello,", name)
}
