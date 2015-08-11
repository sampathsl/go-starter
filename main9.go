package main
import (
	"fmt"
	//"time"
	"sync"
)
type Job struct {
	i int
	max int
	text string
}
func outputText1(j *Job , goGroup *sync.WaitGroup) {
	for j.i < j.max {
		//time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s - %d \n", j.text , j.i )
		j.i++
	}
	goGroup.Done()
}
func main() {
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 2
	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 2
	fmt.Println("1")
	go outputText1(hello, goGroup)
	fmt.Println("2")
	go outputText1(world, goGroup)
	fmt.Println("3")
	goGroup.Add(2)
	goGroup.Wait()
	fmt.Println("Ending")
}