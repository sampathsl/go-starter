package main

type member struct {
	name string
	race string
}

func (m member) intro() {
	println("Hi, I'm " + m.name)
}

func main() {
	krillin := member{
		name: "Krillin",
		race: "Human",
	}

	krillin.intro()
}