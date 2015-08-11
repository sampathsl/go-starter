package main

func main() {
	t3s := map[string]string{
		"Krillin": "Human",
		"Gohan":   "Saiyan",
		"Dende":   "Namekian",
	}

	for name, race := range t3s {
		println(name + " is a " + race)
	}
}