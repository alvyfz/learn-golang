package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {

	alvy := Man{"Alvy Fauzi"}
	alvy.Married()

	fmt.Println(alvy.Name)

}
