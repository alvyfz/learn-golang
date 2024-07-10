package main

func sum(args ...int) {
	total := 0
	for _, v := range args {
		total += v
	}
	println(total)
}

func main() {
	sum(1, 2, 3, 4, 5)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	sum(numbers...) // harus di iterate dulu
}
