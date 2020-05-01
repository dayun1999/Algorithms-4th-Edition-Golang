package SelectionSort

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

//SelectionSort sorts the data which implements the sort.Interface
func SelectionSort(object sort.Interface) {
	length := object.Len()
	for i := 0; i < length-1; i++ {
		min := i
		for j := min + 1; j < length; j++ {
			if object.Less(j, min) {
				min = j
			}
		}
		if min != i {
			object.Swap(min, i)
		}
	}
}
