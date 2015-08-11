package main

import(
	"fmt"
	"sync"
	"runtime"
	"strings"
	"time"
)

var initialString string
var finalString string
var stringLength int

func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	finalString += letter
	wg.Done()
}

func capitalize2(letterChannel chan string, currentLetter string,  wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letterChannel <- thisLetter
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	initialString = `Four score and seven years ago our fathers`+
	`brought forth on this continent , a new nation, conceived in`+
	`Liberty, and dedicated to the proposition that all men are`+
	`created equal.`
	start := time.Now()
	initialBytes := []byte(initialString)
	var letterChannel chan string = make(chan string)
	stringLength = len(initialBytes)
	for i := 0; i < stringLength; i++ {
		wg.Add(2)
		go capitalize2(letterChannel, string(initialBytes[i]), &wg)
		go addToFinalStack(letterChannel, &wg)
		wg.Wait()
	}
	fmt.Println(finalString)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}
