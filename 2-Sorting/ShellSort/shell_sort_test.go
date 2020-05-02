package ShellSort

import (
	"reflect"
	"sort"
	"testing"
)

var testCases = []struct {
	object   sort.Interface
	expected sort.Interface
}{
	{
		StringSlice{"abc", "xyz", "flag"},
		StringSlice{"abc", "flag", "xyz"},
	},
	{
		StringSlice{"me", "you", "us", "hay", "hai", "good", "bad", "linux", "java", "python", "rust", "php"},
		StringSlice{"bad", "good", "hai", "hay", "java", "linux", "me", "php", "python", "rust", "us", "you"},
	},
	{
		StringSlice{"012", "123", "234", "8", "789"},
		StringSlice{"012", "123", "234", "789", "8"},
	},
	{
		StringSlice{},
		StringSlice{},
	},
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

func TestShellSort(t *testing.T) {
	for _, tc := range testCases {
		ShellSort(tc.object)
		if !reflect.DeepEqual(tc.object, tc.expected) {
			t.Errorf("期待的结果是%v,然而得到的却是%v\n", tc.expected, tc.object)
		}
	}
}
