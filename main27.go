package main

import(
	"sync"
	"runtime"
	"fmt"
)

var finStr string

func addToFinalStackX(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	finStr += letter
	wg.Done()
}

func doNothingX(letterChannel chan string, wg *sync.WaitGroup) {
	wg.Done()
	letterChannel <- "nothing"
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var channel chan string = make(chan string)
	wg.Add(2)
	go doNothingX(channel,&wg)
	go addToFinalStackX(channel,&wg)
	wg.Wait()
	fmt.Println(finStr)
}
