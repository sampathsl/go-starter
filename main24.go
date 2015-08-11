package main

import "time"

var greeting string

func init() {
	h := time.Now().Hour()

	switch {
	case h < 12:
		greeting = "Good Morning"
	case h < 18:
		greeting = "Good Evening"
	case h < 24:
		greeting = "Good Night"
	default:
		greeting = "Good Lord!!!"
	}
}

// Print prints a greeting to stdout
func Print1() {
	println(greeting)
}

func main() {
	Print1()
}