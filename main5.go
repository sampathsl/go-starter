package main
import "fmt"
type Artist struct {
	Name, Genre string
	Songs int
}
//Pass by reference ** value change without assigning (pass by value)
func newRelease1(a *Artist) int {
	a.Songs++
	return a.Songs
}
func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease1(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
