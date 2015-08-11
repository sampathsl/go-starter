package main

import "fmt"

func location_(name, city string) (name1, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	name1 = name
	return
}
func main() {
	name, continent := location_("Matt", "LA")
	fmt.Printf("%s lives in %s", name, continent)
}
