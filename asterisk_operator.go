package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"Tasikmalaya", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1, 1) // berubah
	fmt.Println(address2, 2) // Berubah

	// address2 = &Address{"Jakarta", "Jawa Barat", "Indonesia"}
	// fmt.Println(address1, 3) // tetap yang sblmnnya
	// fmt.Println(address2, 4) // akan berubah

	*address2 = Address{"Bogor", "Jawa Barat", "Indonesia"} // Ini akan merubah semua data yang terpointer
	fmt.Println(address1, 5)                                // berubah
	fmt.Println(address2, 6)                                // akan berubah

}
