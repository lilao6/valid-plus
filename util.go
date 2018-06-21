package valid

import (
	"fmt"
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

var (
	// key: function name
	// value: the number of parameters
	funcs = make(Funcs)

	// doesn't belong to validation functions
	unFuncs = map[string]bool{
		"Clear":     true,
		"HasErrors": true,
		"ErrorMap":  true,
		"Error":     true,
		"apply":     true,
		"Check":     true,
		"Valid":     true,
		"NoMatch":   true,
	}
)

func init() {
	v := &Validation{}
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if !unFuncs[m.Name] {
			funcs[m.Name] = m.Func
		}
	}
}

// CustomFunc is for custom validate function
type CustomFunc func(v *Validation, obj interface{}, key string)

// ValidFunc Valid function type
type ValidFunc struct {
	Name   string
	Params []interface{}
}

// Funcs Validate function map
type Funcs map[string]reflect.Value

// Call validate values with named type string
func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	if _, ok := f[name]; !ok {
		err = fmt.Errorf("%s does not exist", name)
		return
	}
	if len(params) != f[name].Type().NumIn() {
		err = fmt.Errorf("The number of params is not adapted")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in)
	return
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
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
