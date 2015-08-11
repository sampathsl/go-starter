package main
import (
	"fmt"
	"time"
)
type Job struct {
	i int
	max int
	text string
}
func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s - %d \n", j.text , j.i )
		j.i++
	}
}
func main() {
	hello := new(Job)
	world := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3
	world.text = "world"
	world.i = 0
	world.max = 5
	fmt.Println(":FUCK")
	go outputText(hello)
	fmt.Println(":FICK")
	outputText(world)
	//go outputText(world)
	fmt.Println("END MAIN")
}