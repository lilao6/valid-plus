/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

var Model bool

const (
	RULE_DEBUG = true
	RULE_RELEASE = false
)

func init() {
	var rules = []Rule{
		new(EmailRule),
		new(InRule),
		new(IpRule),
		new(MaxRule),
		new(MinRule),
		new(NumberRule),
		new(OrRule),
		new(PhoneRule),
		new(LengthRule),
		new(AlphaRule),
		new(Base64Rule),
		new(TelRule),
		new(ZipCodeRule),
		new(CharacterRule),
		new(CharacterNumberRule),
	}
	for _, rule := range rules {
		err := RegisterRule(rule)
		if err != nil {
			panic("register rule:" + err.Error())
			return
		}
	}
	Model = RULE_RELEASE
}
