package utils

import "reflect"

//判断Slice中是否含有某个元素
func IncludeItem(array interface{}, target interface{}) (int, bool) {

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(target, s.Index(i).Interface()) == true {
				return i, true
			}
		}
	}
	return -1, false
}

