/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"errors"
	"regexp"
	"strconv"
)

/*
* <p>移动：134(0-8)、135、136、137、138、139、147、150、151、152、157、158、159、178、182、183、184、187、188、198</p>
* <p>联通：130、131、132、145、155、156、175、176、185、186、166</p>
* <p>电信：133、153、173、177、180、181、189、199</p>
* <p>全球星：1349</p>
* <p>虚拟运营商：170</p>
*/
var phoneRegexp = regexp.MustCompile("^(86|\\+86)?((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3|5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")

// 移动手机号码
type PhoneRule struct {
	value string
	FullTag
}

func (r *PhoneRule) Clone() Rule {
	clone := *r
	return &clone
}

func (r *PhoneRule) Tag() string {
	return "Phone"
}

func (r *PhoneRule) Generate(value interface{}, tagValue string) error {
	if value == nil {
		return errors.New("Generate Phone:value is nil")
	}
	switch value.(type) {
	case string:
		r.value = value.(string)
	case int, int64:
		r.value = strconv.FormatInt(intAll2int(value), 10)
	case uint, uint64:
		r.value = strconv.FormatUint(uintAll2uint(value), 10)
	default:
		return errors.New("Generate Phone:Only number and string data types are supported")
	}
	return nil
}

func (r *PhoneRule) Valid() error {
	if r.value == "" {
		return errors.New("Validation Phone:value is nil")
	}
	if !phoneRegexp.MatchString(r.value) {
		return errors.New("Validation Phone:the string value verification failed")
	}
	return nil
}
