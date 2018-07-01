/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package example

import (
	"github.com/golyu/valid"
	"fmt"
)

func TestRecursive() {
	type A struct {
		Sex int `Must;valid:"Max(2);Min(1);ErrorCode(1111)"` // 最大数值为2,最小数值为1
	}
	type B struct {
		Age int `Must;valid:"Max(23);Min(18);ErrorCode(1111)"` // 最大长度为23,最小长度为18
		A
	}
	v := new(valid.Validation)
	b, code, err := v.Valid(
		&B{
			Age: 15,
			A:   A{Sex: 1},
		})
	fmt.Println(b,code,err)
}
