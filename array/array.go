package array

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/514366607/tools/logger"
	"github.com/514366607/tools/sub"
)

// IN 判断是否在数组里,性能好差的，最好不要用。直接用FormatSliceToMapKey然后判断是否在里面
func IN(data interface{}, dataArray interface{}) bool {
	dataValue := reflect.ValueOf(dataArray)
	if dataValue.Len() == 0 {
		return false
	}

	switch dataValue.Kind() {
	case reflect.Slice:
		for i := 0; i < dataValue.Len(); i++ {
			if data == dataValue.Index(i).Interface() {
				return true
			}
		}
		break
	case reflect.Map:
		for _, key := range dataValue.MapKeys() {
			if data == dataValue.MapIndex(key).Interface() {
				return true
			}
		}
		break
	default:
		logger.Error("传入不为Slice|Map")
	}

	return false
}

// Sum 取得数组和
func Sum(dataArray interface{}) (int64, error) {
	dataValue := reflect.ValueOf(dataArray)
	if dataValue.Len() == 0 {
		return 0, errors.New("数组为空")
	}

	var sum int64
	switch dataValue.Kind() {
	case reflect.Slice:
		for i := 0; i < dataValue.Len(); i++ {
			switch dataValue.Index(i).Kind() {
			case reflect.Int:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(int)))
			case reflect.Int8:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(int8)))
			case reflect.Int16:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(int16)))
			case reflect.Int32:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(int32)))
			case reflect.Int64:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(int64)))
			case reflect.Uint:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(uint)))
			case reflect.Uint8:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(uint8)))
			case reflect.Uint16:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(uint16)))
			case reflect.Uint32:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(uint32)))
			case reflect.Uint64:
				sum = sub.Int64(sum, int64(dataValue.Index(i).Interface().(uint64)))
			default:
				return 0, fmt.Errorf("错误类型%v", dataValue.Index(i).Kind())
			}
		}
		break
	case reflect.Map:
		for _, key := range dataValue.MapKeys() {
			switch dataValue.MapIndex(key).Kind() {
			case reflect.Int:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(int)))
			case reflect.Int8:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(int8)))
			case reflect.Int16:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(int16)))
			case reflect.Int32:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(int32)))
			case reflect.Int64:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(int64)))
			case reflect.Uint:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(uint)))
			case reflect.Uint8:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(uint8)))
			case reflect.Uint16:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(uint16)))
			case reflect.Uint32:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(uint32)))
			case reflect.Uint64:
				sum = sub.Int64(sum, int64(dataValue.MapIndex(key).Interface().(uint64)))
			default:
				return 0, fmt.Errorf("错误类型%v", dataValue.MapIndex(key).Kind())
			}
		}
		break
	default:
		logger.Error("传入不为Slice|Map")
	}
	return sum, nil
}
