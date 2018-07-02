/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package example

import (
	"github.com/golyu/valid/rule"
	"errors"
	"regexp"
	"github.com/golyu/valid"
	"fmt"
)

func customRule() {
	qqRule := new(QQRule)
	rule.RegisterRule(qqRule)

	type Xx struct {
		MyQQ string `valid:"Must;QQ;ErrorCode(1111)"`
	}
	v := new(valid.Validation)
	fmt.Println(v.Valid(&Xx{MyQQ: "2131121212"}))
}

// qq号码校验规则
type QQRule struct {
	value string
	rule.FullTag
}

func (r *QQRule) Clone() rule.Rule {
	clone := *r
	return &clone
}

var qqRegexp = regexp.MustCompile("[1-9][0-9]{4,14}")

func (*QQRule) Tag() string {
	return "QQ"
}

func (r *QQRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate QQ:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate QQ:the value generate failed")
	}
	return nil
}

func (r *QQRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation QQ:value is nil")
	}
	if !qqRegexp.MatchString(r.value) {
		return errors.New("Validation QQ:the string value verification failed")
	}
	return nil
}
