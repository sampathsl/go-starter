
package main

import "fmt"

func main() {
	sz := 3 // slice size
	t3s := make([]string, sz)
	t3s[0] = "Krillin"
	t3s[1] = "Gohan"
	t3s[2] = "Dende"
	fmt.Println(t3s)
	t3ss := []string{"Krillin", "Gohan", "Dende"}
	t4s := append(t3ss, "Vegeta")
	fmt.Println(t3ss, t4s)
	copy(t3s, t4s)
	fmt.Println(t3s, t4s)
	fmt.Println("----------------------------------------")
	t4s1 := []string{"Krillin", "Gohan", "Dende", "Vegeta"}
	t3s1 := make([]string, 3)
	copy(t3s1, t4s1)
	fmt.Println(t3s1,t4s1)
}