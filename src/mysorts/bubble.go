package mysorts

import "fmt"

func BubbleSort(list []int) []int {
	swaps := 0
	for index := range list {
		i := index
		for j := index + 1; j < len(list); j++ {
			if list[j] < list[i] {
				swap(i, j, list)
				swaps++
			}
			i++
		}
	}
	fmt.Println(list)
	fmt.Printf("%d swaps\n", swaps)
	if swaps > 0 {
		return BubbleSort(list)
	}
	return list
}
