package main

import "fmt"

func endApp() {
	fmt.Println("Application berakhir")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApplication(error bool) {
	defer endApp()

	if error {
		panic("Terjadi error")
	}

}

func main() {
	runApplication(true)
	fmt.Println("Recover")
}
