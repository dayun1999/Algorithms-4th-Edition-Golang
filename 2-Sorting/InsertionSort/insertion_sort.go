package InsertionSort

import "sort"

type IntegerSlice []int

func (l IntegerSlice) Len() int {
	return len(l)
}

func (l IntegerSlice) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l IntegerSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//InsertionSort takes any variable which implements the sort.Interface
func InsertionSort(object sort.Interface) {
	length := object.Len()
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		for j := i; j > 0 && object.Less(j, j-1); j-- {
			object.Swap(j, j-1)
		}
	}
}
