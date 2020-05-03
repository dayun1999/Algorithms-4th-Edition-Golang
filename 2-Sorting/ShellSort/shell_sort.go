package ShellSort

import "sort"

type StringSlice []string
type IntegerSlice []int

//ShellSort sorts the data which implements the sort.Interface
func ShellSort(object sort.Interface) {
	n := object.Len()
	if n <= 1 {
		return
	}
	// 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
	//一般，步长序列(增长序列)不同，该算法的(最坏)时间复杂度也不尽相同，想深入研究请移步维基百科
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && object.Less(j, j-h); j -= h {
				object.Swap(j, j-h)
			}
		}
		h /= 3
	}
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	//object[i], object[j] = object[j], object[i]
}

func (s StringSlice) Len() int {
	return len(s)
}

func (s IntegerSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntegerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	//object[i], object[j] = object[j], object[i]
}

func (s IntegerSlice) Len() int {
	return len(s)
}

//如果不实现sort包的接口(实现接口是为了拓展性更强，如果我没写错的话)，一般的实现是这样的
//func ShellSort(object sort.Interface) {
//	n := len(object)
//	if n <= 1 {
//		return
//	}
////希尔排序步长的选取直接影响算法的复杂度
//	h := 1
//	for h < n/3 {
//		h = 3*h + 1
//	}
//	for h >= 1 {
//		for i := h; i < n; i++ {
//			for j := i; j >=h && less(object[j], object[j-h]); j -= h {
//				swap(object, j, j-h)
//			}
//		}
//		h /= 3
//	}
//}
