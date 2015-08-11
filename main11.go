package main
import (
	"fmt"
	"time"
	"io/ioutil"
)
type Job struct {
	i int
	max int
	text string
}
func outputText2(j *Job) {
	fileName := j.text + ".txt"
	fileContents := ""
	fmt.Println("XXX1")
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	fmt.Println("XXX2")
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if (err != nil) {
		fmt.Println("err" + err.Error())
		panic("Something went awry")
	}
}
func main() {
	fmt.Println("Starting")
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3
	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 5
	fmt.Println("1")
	go outputText2(hello)
	fmt.Println("2")
	go outputText2(world)
	fmt.Println("3")
	fmt.Println("Ending")
}