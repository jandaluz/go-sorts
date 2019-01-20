package mysorts

import (
	"fmt"
	"mysorts"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	copyIntList := make([]int, len(intList))
	copy(copyIntList, intList)
	copyIntList = mysorts.InsertionSort(copyIntList)
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
	copyIntList = mysorts.InsertionSort(copyIntList)
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
	copyIntList = mysorts.InsertionSort(copyIntList)
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
