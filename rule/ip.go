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

// ip正则表达式
var ipRegexp = regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)

type IpRule struct {
	value string
	FullTag
}

func (r *IpRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *IpRule) Tag() string {
	return "IP"
}

func (r *IpRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Email:value is nil")
	}
	var ok bool
	r.value, ok = value.(string)
	if !ok {
		return errors.New("Generate IP:the value generate failed")
	}
	return nil
}

func (r *IpRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation IP:value is nil")
	}
	if !ipRegexp.MatchString(r.value) {
		return errors.New("Validation IP:the int value verification failed")
	}
	return nil
}
