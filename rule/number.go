/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"errors"
	"strconv"
)

// 数字校验
type NumberRule struct {
	value string
	FullTag
}

func (r *NumberRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *NumberRule) Tag() string {
	return "Number"
}

func (r *NumberRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Number:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate Number:the value generate failed")
	}
	return nil
}

func (r *NumberRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation Number:value is nil")
	}
	_, err := strconv.Atoi(r.value)
	if err != nil {
		return errors.New("Validation Number:the string value verification failed")
	}
	return nil
}
