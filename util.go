package valid

import (
	"reflect"
)

const (
	// 校验框架的tag
	VALID_TAG = "valid"
	// 是否必传
	MUST = "Must"
	// 分割各个规则的分隔符
	SPLIT_SEP = ";"
)

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isSlice(t reflect.Type) bool {
	return t.Kind() == reflect.Slice
}

func isSlicePtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Slice
}

func isPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

// IsReflectZeroValue 是否零值
func IsReflectZeroValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		temp := value.Int()
		return temp == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		temp := value.Uint()
		return temp == 0
	case reflect.Float32, reflect.Float64:
		temp := value.Float()
		return temp == 0
	case reflect.String:
		return value.String() == ""
	case reflect.Slice:
		return value.Len() == 0
	default:
		return true
	}
}
