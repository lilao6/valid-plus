/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"sync"
	"errors"
	"strings"
)

var (
	rules sync.Map
)

// RegisterRule 注册一个规则
func RegisterRule(rule Rule) error {
	if rule == nil {
		return errors.New("Register rule failure:rule is nil")
	}
	_, ok := rules.Load(rule.Tag())
	if ok {
		return errors.New("Register rule failure:Rule has been registered")
	}
	rules.Store(rule.Tag(), rule)
	return nil
}

// GetRule 获取一个规则
func GetRule(fullTag string) (Rule, error) {
	if fullTag == "" {
		return nil, errors.New("Get rule failure:tag is nil")
	}
	var tag string
	// Or<Max(1)>    -->   Or
	if i := strings.Index(fullTag, TAG_PRE_OR); i > 0 {
		tag = fullTag[:i]
		// Max(1)     -->   max
	} else if i := strings.Index(fullTag, TAG_PRE); i > 0 {
		tag = fullTag[:i]
	} else {
		tag = fullTag
	}
	var rule Rule
	temp, ok := rules.Load(tag)
	if !ok {
		return nil, errors.New("Get rule failure:The <" + fullTag + "> rule  does not exist")
	}
	rule = temp.(Rule).Clone()
	rule.SetFullTag(fullTag)
	return rule, nil
}
