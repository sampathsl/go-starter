package main

import "fmt"

func location(city string) (string, string , int) {
	var region string
	var continent string
	var id int
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent , id = "California", "North America" , 1
	case "New York", "NYC":
		region, continent , id = "New York", "North America" , 2
	default:
		region, continent , id = "Unknown", "Unknown" , 3
	}
	return region, continent , id
}

func main() {
	region, continent , id := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s , %d", region, continent , id)
}
