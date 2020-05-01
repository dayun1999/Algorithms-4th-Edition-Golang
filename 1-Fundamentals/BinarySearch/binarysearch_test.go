package BinarySearch

import "testing"

var testCases = []struct {
	object   []int
	target   int
	expected int
}{
	{
		[]int{20, 23, 44, 55, 89, 90, 100, 101},
		101,
		7,
	},
	{
		[]int{1, 2, 5, 78, 100, 108, 134},
		2,
		1,
	},
	{
		[]int{0, 1, 2, 3, 5, 9, 10, 11},
		8,
		-1,
	},
	{
		[]int{},
		90,
		-1,
	},
	{
		[]int{1, 2},
		1,
		0,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range testCases {
		index := BinarySearch(tc.object, tc.target)
		if index != tc.expected {
			t.Errorf("期待的索引是%d,得到的却是%d\n", tc.expected, index)
		}
	}
	t.Logf("测试完成\n")
}
