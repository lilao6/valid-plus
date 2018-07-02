/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"regexp"
	"errors"
)

var emailRegexp = regexp.MustCompile("^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")

// 邮箱校验
type EmailRule struct {
	value string
	FullTag
}

func (r *EmailRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *EmailRule) Tag() string {
	return "Email"
}

func (r *EmailRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Email:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate Email:the value generate failed")
	}
	return nil
}

func (r *EmailRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation Email:value is nil")
	}
	if !emailRegexp.MatchString(r.value) {
		return errors.New("Validation Email:the string value verification failed")
	}
	return nil
}
