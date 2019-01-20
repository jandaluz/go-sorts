package mysorts

import "fmt"

func InsertionSort(list []int) []int {

	sorted := list[0:1]
	swaps := 0
	//for i := 2; i < len(list)+1; i++ {
	for i := 1; i < len(list); i++ {
		sorted = list[0 : i+1]
		for j := i; j > 0; j-- {
			fmt.Println(sorted[j], sorted[j-1])
			if sorted[j] < sorted[j-1] {
				fmt.Println(sorted[j], sorted[j-1])
				swap(j, j-1, sorted)
				swaps++
			}
		}
		fmt.Println(sorted)
		fmt.Println("SWAPS:", swaps)

	}
	return sorted
}
