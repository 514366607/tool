package array

import "reflect"

// FormatSliceToMapKey 将Slice转为Map返回方便确认是否在数组里
func FormatSliceToMapKey(iDs interface{}) map[interface{}]struct{} {
	v := reflect.ValueOf(iDs)
	if v.Kind() != reflect.Slice {
		panic("参数错误，不为Slice")
	}

	var exIDsMap = make(map[interface{}]struct{})
	for i := 0; i < v.Len(); i++ {
		exIDsMap[v.Index(i).Interface()] = struct{}{}
	}
	return exIDsMap
}

// InInt64Slice 判断一个元素是否在int64切片里面
func InInt64Slice(element int64, slice []int64) bool {
	for _, vv := range slice {
		if vv == element {
			return true
		}
	}
	return false
}

// InIntSlice 判断一个元素是否在int切片里面
func InIntSlice(element int, slice []int) bool {
	for _, vv := range slice {
		if vv == element {
			return true
		}
	}
	return false
}

// InInt16Slice 判断一个元素是否在int16切片里面
func InInt16Slice(element int16, slice []int16) bool {
	for _, vv := range slice {
		if vv == element {
			return true
		}
	}
	return false
}
