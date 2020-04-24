package utils

import "sort"

// way != -1 正序, 默认正序
func SortStr(str []string, way int) []string {
	if way != -1 {
		//正序
		return sortStr(str)
	}
	return reverse(str)
}

func reverse(str []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	return str
}
func sortStr(str []string) []string {
	sort.Strings(str)
	return str
}
