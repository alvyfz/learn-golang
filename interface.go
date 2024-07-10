package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {

	alvy := Person{Name: "Alvy Fauzi"}

	SayHello(alvy)

}
