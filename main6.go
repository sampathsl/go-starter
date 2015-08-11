// We have to count some fruits using this simple go program.
// The total fruit count should be 12000000 but the result is incorrect.
// Find a good way to solve this problem. And use the "Share" button above
// to get a link to your go code. Give the link as the answer.
// You must count fruits one by one (eg count function).

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	gophers = 4
)

var (
	apples  int64 = 2000000
	mangos  int64 = 4000000
	bananas int64 = 5000000
	oranges int64 = 1000000
	fruits  int64 = 0
	wg sync.WaitGroup
	l sync.Mutex
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(gophers)
	wg.Add(4)

	go count(apples)
	go count(mangos)
	go count(bananas)
	go count(oranges)

	wg.Wait()
	fmt.Printf("There are %d fruits\n", fruits)
	if fruits != 12000000 {
		fmt.Printf("THERE SHOULD BE %d FRUITS\n", 12000000)
	}
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}

func count(f int64) {
	defer wg.Done()
	l.Lock()
	for f > 0 {
		f -= 1
		fruits++
	}
	l.Unlock()
}
