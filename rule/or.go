/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"errors"
	"strings"
)

// 满足一个条件就校验成功了,因为Or这个条件控制比较特殊,所以跟其他的不一样,使用Or<>包裹条件,各条件之间用空格符隔开
// 例如,想要一个联系方式的字符串校验,这个联系方式可以是电话号码Phone,也可以是email,就可以写为 Or<Phone Email>
type OrRule struct {
	//childRules []Rule
	childTags []string
	value     interface{}
	FullTag
}

func (r *OrRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *OrRule) Tag() string {
	return "Or"
}

func (r *OrRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Or:value is nil")
	}
	tagValueLen := len(tagValue)
	tagLen := len(r.Tag())
	if tagValueLen <= tagLen+2 { // 2的长度是()两个括号
		return errors.New("Generate Or:the tag are out of specification")
	}
	// Or<Phone Max(3)>  -->   []string{"Phone","max(3)"}
	tagValues := strings.Split(tagValue[tagLen+1:tagValueLen-1], SPLIT_SEP_OR)
	r.value = value
	r.childTags = tagValues
	return nil
}

func (r *OrRule) Valid() error {
	var probe bool // 探针,只要一次成功就认为是成功了
	var errTag []string
	for _, tag := range r.childTags {
		// 获取指定规则
		rule, err := GetRule(tag)
		if err != nil {
			return errors.New("Generate Or:" + err.Error())
		}
		// 生成规则数据
		err = rule.Generate(r.value, tag)
		if err != nil {
			return errors.New("Generate Or:" + err.Error())
		}
		// 校验规则
		err = rule.Valid()
		if err == nil {
			probe = true
		}
		errTag = append(errTag, rule.Tag())
	}
	if probe {
		return nil
	}
	return errors.New("Validation Or:the <" + strings.Join(errTag, ",") + "> rule verification failed")
}
