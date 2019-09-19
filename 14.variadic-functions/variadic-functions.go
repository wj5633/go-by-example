package main

import "fmt"

func sun(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sun(1, 2)
	sun(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	sun(nums...)
}
