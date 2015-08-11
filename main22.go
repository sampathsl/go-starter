package main

import "fmt"

func main() {
	type member struct {
		name string
		race string
	}

	krillin := member{}
	krillin.name = "Krillin"
	krillin.race = "Human"

	gohan := member{
		name: "Gohan",
		race: "Saiyan",
	}

	fmt.Println(krillin, gohan)
}