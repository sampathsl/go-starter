package main
import(
	"runtime"
	"fmt"
	"time"
	"strconv"
)
func showNumber1(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
}
func main() {
	//runtime.GOMAXPROCS(4)
	iterations := 10
	for i := 0 ; i <= iterations ; i++ {
		go showNumber1(i)
	}
	runtime.Gosched()
	fmt.Println("Goodbye!")
}
