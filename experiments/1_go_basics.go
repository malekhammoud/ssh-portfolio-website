package main

import "fmt"

var skills = [3] string {"Python", "JS", "GO"}
func main() {
	var name = "Malek Hammoud."
	fmt.Println("hello.")
	fmt.Println("Welcome.")
	fmt.Printf("I'm %s \n", name)
	fmt.Println("I Build Software that has to work")

	var userName string
	fmt.Println("What's your name?")
	fmt.Print(">")
	fmt.Scan(&userName)
	fmt.Printf("hello %s \n", userName)

	for _, skill := range skills{
		fmt.Println(skill)
	}
}

