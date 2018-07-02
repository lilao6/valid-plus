/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"errors"
	"unicode/utf8"
	"reflect"
	"strconv"
)

// 校验最小值和长度
type MinRule struct {
	value    interface{}
	min      int64   // 作用于int,string,slice类型
	uMin     uint64  // 作用于uint类型
	minFloat float64 // 作用于float类型
	FullTag
}

func (r *MinRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *MinRule) Tag() string {
	return "Min"
}

func (r *MinRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Min:value is nil")
	}
	r.value = value
	tagValueLen := len(tagValue)
	tagLen := len(r.Tag())
	if tagValueLen <= tagLen+2 { // 2的长度是()两个括号
		return errors.New("Generate Min:the tag are out of specification")
	}
	// Min(1)  -->   "1"
	tagValue = tagValue[tagLen+1 : tagValueLen-1]
	switch r.value.(type) {
	case int, int8, int16, int32, int64:
		var err error
		r.min, err = strconv.ParseInt(tagValue, 10, 64)
		if err != nil {
			return errors.New("Generate Min:" + err.Error())
		}
	case float32, float64:
		var err error
		r.minFloat, err = strconv.ParseFloat(tagValue, 64)
		if err != nil {
			return errors.New("Generate Min:" + err.Error())
		}
	case uint, uint8, uint16, uint32, uint64, string, []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []string:
		var err error
		r.uMin, err = strconv.ParseUint(tagValue, 10, 64)
		if err != nil {
			return errors.New("Generate Min:" + err.Error())
		}
	default:
		return errors.New("Generate Min:Only basic and slice data types are supported")
	}
	return nil
}

func (r *MinRule) Valid() error {
	if r.value == nil {
		return errors.New("Validation Min:value is nil")
	}
	switch r.value.(type) {
	case int, int8, int16, int32, int64:
		value := intAll2int(r.value)
		if value < r.min {
			return errors.New("Validation Min:the int value verification failed")
		}
	case uint, uint8, uint16, uint32, uint64:
		value := uintAll2uint(r.value)
		if value < r.uMin {
			return errors.New("Validation Min:the uint value verification failed")
		}
	case float32, float64:
		value := floatAll2float(r.value)
		// 考虑精度误差
		if r.minFloat-value > ACCURACY {
			return errors.New("Validation Min:the float value verification failed")
		}
	case string:
		value := r.value.(string)
		length := utf8.RuneCountInString(value)
		if uint64(length) < r.uMin {
			return errors.New("Validation Min:the string value length verification failed")
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []string:
		length := reflect.ValueOf(r.value).Len()
		if int64(length) < r.min {
			return errors.New("Validation Min:the slice value length verification failed")
		}
	}
	return nil
}
