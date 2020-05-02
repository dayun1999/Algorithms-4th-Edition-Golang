package InsertionSort

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	object   IntegerSlice
	expected IntegerSlice
}{
	{
		IntegerSlice{101, 100, 90, 89, 55, 44, 23, 20},
		IntegerSlice{20, 23, 44, 55, 89, 90, 100, 101},
	},
	{
		IntegerSlice{5, 2, 1, 78, 108, 134, 100},
		IntegerSlice{1, 2, 5, 78, 100, 108, 134},
	},
	{
		IntegerSlice{1, 2, 10, 11, 9, 5, 3, 0},
		IntegerSlice{0, 1, 2, 3, 5, 9, 10, 11},
	},
	{
		IntegerSlice{9, 8, 90, 120, 67676, 23, 56, -1, -90},
		IntegerSlice{-90, -1, 8, 9, 23, 56, 90, 120, 67676},
	},
	{
		IntegerSlice{},
		IntegerSlice{},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		InsertionSort(tc.object)
		if !reflect.DeepEqual(tc.object, tc.expected) {
			t.Errorf("期望得到的是%v,然而得到的却是%v\n", tc.expected, tc.object)
		}
	}
}
