package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var alvy Customer
	fmt.Println(alvy)
	alvy.Name = "Muhammad Alvy Eka Fauzi"
	alvy.Address = "Tasikmalaya"
	alvy.Age = 23

	fmt.Println(alvy)
	fmt.Println(alvy.Name)
	fmt.Println(alvy.Address)
	fmt.Println(alvy.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(joko)
	alex := Customer{
		"Alex",
		"Jakarta",
		30,
	}
	fmt.Println(alex)
	alex.sayHello("Alvy")
}
