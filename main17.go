package main
import (
	"fmt"
	"runtime"
	"strings"
	"time"
)
var loremIpsum string
var finalIpsum string
//var letterSentChan chan string
func deliverToFinal1(letter string, finalIpsum *string) {
	*finalIpsum += letter
}
func capitalize1(current *int, length int, letters []byte, finalIpsum *string) {
	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal1(thisLetter, finalIpsum)
		*current++
	}
}
func main() {
	start := time.Now()
	runtime.GOMAXPROCS(8)
	index := new(int)
	*index = 0
	loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing`+
	`elit. Vestibulum venenatis magna eget libero tincidunt, ac`+
	`condimentum enim auctor. Integer mauris arcu, dignissim sit amet`+
	`convallis vitae, ornare vel odio. Phasellus in lectus risus. Ut`+
	`sodales vehicula ligula eu ultricies. Fusce vulputate fringilla`+
	`eros at congue. Nulla tempor neque enim, non malesuada arcu`+
	`laoreet quis. Aliquam eget magna metus. Vivamus lacinia`+
	`venenatis dolor, blandit faucibus mi iaculis quis. Vestibulum`+
	`sit amet feugiat ante, eu porta justo.`
	letters := []byte(loremIpsum)
	length := len(letters)
	go capitalize1(index, length, letters, &finalIpsum)
	/*go func() {
		go capitalize(index, length, letters, &finalIpsum)
	}()*/
	fmt.Println("Total of " , length , " characters.")
	fmt.Println("loremIpsum := "+loremIpsum)
	fmt.Println("index (Loop amount) := " , *index)
	fmt.Println("finalIpsum (Converted Letters) := " , finalIpsum)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}