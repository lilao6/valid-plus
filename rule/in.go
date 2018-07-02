/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"errors"
	"math"
	"strings"
)

// 校验数据是否存在于所列条件之中
type InRule struct {
	value        interface{}
	intValues    []int64
	uintValues   []uint64
	floatValues  []float64
	stringValues []string
	FullTag
}

func (r *InRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *InRule) Tag() string {
	return "In"
}

func (r *InRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate In:value is nil")
	}
	r.value = value
	tagValueLen := len(tagValue)
	tagLen := len(r.Tag())
	if tagValueLen <= tagLen+2 { // 2的长度是()两个括号
		return errors.New("Generate In:the tag are out of specification")
	}
	// in(1,2,3)  -->   []string{"1","2","3"}
	tagValues := strings.Split(tagValue[tagLen+1:tagValueLen-1], SPLIT_SEP)
	switch value.(type) {
	case string:
		r.stringValues = tagValues
	case int, int8, int16, int32, int64:
		var err error
		r.intValues, err = strings2Int64s(tagValues)
		if err != nil {
			return errors.New("Generate In:" + err.Error())
		}
	case uint, uint8, uint16, uint32, uint64:
		var err error
		r.uintValues, err = strings2Uint64s(tagValues)
		if err != nil {
			return errors.New("Generate In:" + err.Error())
		}
	case float32, float64:
		var err error
		r.floatValues, err = strings2float64s(tagValues)
		if err != nil {
			return errors.New("Generate In:" + err.Error())
		}
	default:
		return errors.New("Generate In:Only basic data types are supported")
	}

	//var ok bool
	//r.value, ok = value.(string)
	//if !ok {
	//	return errors.New("Generate In:the value generate failed")
	//}
	return nil
}

func (r *InRule) Valid() error {
	if r.value == nil {
		return errors.New("Validation In:value is nil")
	}
	switch r.value.(type) {
	case int, int8, int16, int32, int64:
		value := intAll2int(r.value)
		if r.intValues != nil {
			var probe bool
			for _, temp := range r.intValues {
				if temp == value {
					probe = true
					break
				}
			}
			if !probe {
				return errors.New("Validation In:the int value verification failed")
			}
		}
	case uint, uint8, uint16, uint32, uint64:
		value := uintAll2uint(r.value)
		if r.uintValues != nil {
			var probe bool
			for _, temp := range r.uintValues {
				if temp == value {
					probe = true
					break
				}
			}
			if !probe {
				return errors.New("Validation In:the uint value verification failed")
			}
		}
	case float32, float64:
		value := floatAll2float(r.value)
		if r.floatValues != nil {
			var probe bool
			for _, temp := range r.floatValues {
				if math.Abs(temp-value) < ACCURACY {
					probe = true
					break
				}
			}
			if !probe {
				return errors.New("Validation In:the float value verification failed")
			}
		}
	case string:
		value := r.value.(string)
		if r.stringValues != nil {
			var probe bool
			for _, temp := range r.stringValues {
				if temp == value {
					probe = true
					break
				}
			}
			if !probe {
				return errors.New("Validation In:the string value verification failed")
			}
		}
	}
	return nil
}
