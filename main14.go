package main
import(
	"runtime"
	"fmt"
	"time"
	//"strconv"
)
func showNumber2(num int) {
	//tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	//fmt.Println(num, tstamp)
	fmt.Println(num)
	time.Sleep(time.Millisecond * 1000) // DOES NOT DELAY
}
func main() {
	runtime.GOMAXPROCS(1)
	iterations := 10
	for i := 0 ; i <= iterations ; i++ {
		go showNumber2(i)
		//time.Sleep(time.Millisecond * 1000) // DOES DELAY
	}
	runtime.Gosched()
	fmt.Println("Goodbye!")
}
