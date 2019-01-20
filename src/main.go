package main

import (
	"fmt"
	"mysorts"
)

func main() {
	fmt.Println("Sorts demo")
	intList := []int{5, 4, 3, 2, 1}
	fmt.Println(mysorts.SelecitonSort(intList))

	intList2 := []int{5, 4, 3, 2, 1}
	fmt.Println(mysorts.BubbleSort(intList2))

	intList3 := []int{5, 4, 3, 2, 1}
	fmt.Println(mysorts.InsertionSort(intList3))

	intList4 := []int{22, 5, 9, 11, 15, 30, 99, 159, 55, 63, 44, 330}
	fmt.Println(mysorts.InsertionSort(intList4))
}
