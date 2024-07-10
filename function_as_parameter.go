package main

import "fmt"

type FilterType func(string) string

func sayHelloWithFilter(name string, filter FilterType) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func filterName(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", filterName)
	sayHelloWithFilter("Alvy", filterName)
}
