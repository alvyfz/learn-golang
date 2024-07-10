package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Tasikmalaya", "Jawa Barat", "Indonesia"}
	// address2 := address1 // di copy value nya

	// address2.City = "Bandung"
	// fmt.Println(address1) // Tidak berubah
	// fmt.Println(address2) // Berubah

	var address1 Address = Address{"Tasikmalaya", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // skrng reference

	address2.City = "Bandung"
	fmt.Println(address1) // berubah
	fmt.Println(address2) // Berubah
}
