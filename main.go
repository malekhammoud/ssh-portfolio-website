package main

import "fmt"

func main() {
	var name = "Malek Hammoud."
	fmt.Println("Welcome.")
	fmt.Printf("I'm %s \n", name)
	fmt.Println("I Build Software that has to work")


	var userName string
	fmt.Println("What's your name?")
	fmt.Print(">")
	fmt.Scan(&userName)
	fmt.Printf("hello %s \n", userName)
}
