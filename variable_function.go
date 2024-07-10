package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye    // variable a function
	result := sayGoodBye("Cia") // variable result of function
	fmt.Println(result)
}
