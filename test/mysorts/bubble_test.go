package mysorts

import (
	"fmt"
	"mysorts"
	"sort"
	"testing"
)

var intList = []int{5, 4, 3, 2, 1}
var intList2 = []int{22, 5, 9, 11, 15, 30, 99, 159, 55, 63, 44, 330}
var intList3 = []int{1, 50, 234, 56, 34, 76, 4, 4358, 42, 66, 477, 132, 10, 18, 26, 97, 31, 46, 2455}

func TestBubbleSort(t *testing.T) {
	fmt.Println("inside test bubble sort")
	copyIntList := make([]int, len(intList))
	copy(copyIntList, intList)
	copyIntList = mysorts.BubbleSort(copyIntList)
	fmt.Println("Test Bubble Sort done")
	sorted := sort.SliceIsSorted(copyIntList, func(i, j int) bool {
		fmt.Println(i, j)
		fmt.Println("vals:", copyIntList[i], copyIntList[j])
		return copyIntList[i] < copyIntList[j]
	})

	if !sorted {
		t.Errorf("List was not sorted: %v", copyIntList)
	}
	copyIntList = make([]int, len(intList2))
	copy(copyIntList, intList2)
	fmt.Println(copyIntList)
	copyIntList = mysorts.BubbleSort(copyIntList)
	sorted = sort.SliceIsSorted(copyIntList, func(i, j int) bool {
		if j < len(copyIntList) {
			return copyIntList[i] < copyIntList[j]
		}
		return true
	})
	if !sorted {
		t.Errorf("List was not sorted: %v", copyIntList)
	}

	copyIntList = make([]int, len(intList3))
	copy(copyIntList, intList3)
	fmt.Println("Before: ", copyIntList)
	copyIntList = mysorts.BubbleSort(copyIntList)
	fmt.Println("After: ", copyIntList)
	sorted = sort.SliceIsSorted(copyIntList, func(i, j int) bool {
		if j < len(copyIntList) {
			return copyIntList[i] < copyIntList[j]
		}
		return true
	})
	if !sorted {
		t.Errorf("List was not sorted: %v", copyIntList)
	}

}
