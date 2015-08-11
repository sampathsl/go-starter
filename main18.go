package main
import "fmt"

func main() {
	greet := getGreeter("Hi")
	greet("John Doe") //this is getGreeter's parameter lol :)
}

func getGreeter(prefix string) func(name string) {
	fn := func(name string) {
		fmt.Println("name:= " + name)
		fmt.Println("prefix:= " +prefix)
		println(prefix + " " + name)
	}
	return fn
}