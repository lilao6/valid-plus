package rule

import (
	"errors"
	"strconv"
	"unicode/utf8"
	"reflect"
)

type LengthRule struct {
	value  interface{}
	length int
	FullTag
}

func (r *LengthRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *LengthRule) Tag() string {
	return "Length"
}
func (r *LengthRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Length:value is nil")
	}
	tagValueLen := len(tagValue)
	tagLen := len(r.Tag())
	if tagValueLen <= tagLen+2 { // 2的长度是()两个括号
		return errors.New("Generate Length:the tag are out of specification")
	}
	tagValue = tagValue[tagLen+1 : tagValueLen-1]
	r.value = value
	//tag数值获取
	var err error
	r.length, err = strconv.Atoi(tagValue)
	if nil != err {
		return errors.New("Generate Length:" + err.Error())
	}
	return nil
}

func (r *LengthRule) Valid() error {
	if r.value == nil {
		return errors.New("Validation Length:value is nil")
	}
	//字符串长度匹配
	if str, ok := r.value.(string); ok && (utf8.RuneCountInString(str) == r.length) {
		return nil
	}
	//切片长度匹配
	v := reflect.ValueOf(r.value)
	if v.Kind() == reflect.Slice && v.Len() == r.length {
		return nil
	}
	return errors.New("Validatio length:An array or string is not required for a given length.")
}
