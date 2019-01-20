package mysorts

import "fmt"

func SelecitonSort(list []int) []int {
	for index := range list {
		fmt.Printf("%d\n", index)
		minIndex := index
		for ii := index + 1; ii < len(list); ii++ {
			if list[ii] < list[minIndex] {
				minIndex = ii
			}
		}
		if minIndex != index {
			swap(index, minIndex, list)
		}
	}
	return list
}

func swap(i int, j int, list []int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}
