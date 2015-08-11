
package main

import "fmt"

func main() {

	t3s := make(map[string]string)
	t3s["Krillin"] = "Human"
	t3s["Gohan"] = "Saiyan"
	t3s["Dende"] = "Namekian"
	t3s["Dende"] = "Namekian"
	//fmt.Println(t3s)

	//----------------------

	t3s1 := map[string]string{
		"Krillin": "Human",
		"Gohan": "Saiyan",
		"Dende": "Namekian",
	}

	fmt.Println("XX:" , t3s1)
	t4s1 := t3s1
	t4s1["Vegeta"] = "Saiyan"

	fmt.Println(t3s1)
	fmt.Println(t4s1)

}