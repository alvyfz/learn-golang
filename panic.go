package main

import "fmt"

func endApp() {
	fmt.Println("Application berakhir")
}

func runApplication(error bool) {
	defer endApp()

	if error {
		panic("Terjadi error")
	}
}

func main() {
	runApplication(true)
}
